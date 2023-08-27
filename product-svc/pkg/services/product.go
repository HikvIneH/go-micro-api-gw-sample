package services

import (
	"context"
	"net/http"

	"github.com/hikvineh/product-svc/pkg/db"
	"github.com/hikvineh/product-svc/pkg/models"
	"github.com/hikvineh/product-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) FindAll(context.Context, *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	var products []models.Product

	if result := s.H.DB.Find(&products); result.Error != nil {
		return nil, result.Error
	}

	var dataItems []*pb.FindOneData
	for _, product := range products {
		dataItems = append(dataItems, &pb.FindOneData{
			Id:    product.Id,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		})
	}

	return &pb.FindAllResponse{
		Status: http.StatusOK,
		Data: &pb.FindAllData{
			Items: dataItems,
		},
	}, nil
}


func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1

	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
