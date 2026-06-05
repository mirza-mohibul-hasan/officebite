package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/officebite/officebite/apps/api/internal/models"
	"github.com/officebite/officebite/apps/api/internal/repository"
)

var ErrInvalidMenu = errors.New("invalid menu")

type MenuService struct {
	menus repository.MenuRepository
}

type MenuInput struct {
	Title         string
	Description   string
	Price         int64
	AvailableDate time.Time
}

func NewMenuService(menus repository.MenuRepository) MenuService {
	return MenuService{menus: menus}
}

func (s MenuService) Create(ctx context.Context, input MenuInput) (*models.Menu, error) {
	menu, err := buildMenu(input)
	if err != nil {
		return nil, err
	}

	if err := s.menus.Create(ctx, menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func (s MenuService) Update(ctx context.Context, id uint, input MenuInput) (*models.Menu, error) {
	menu, err := s.menus.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	next, err := buildMenu(input)
	if err != nil {
		return nil, err
	}

	menu.Title = next.Title
	menu.Description = next.Description
	menu.Price = next.Price
	menu.AvailableDate = next.AvailableDate

	if err := s.menus.Update(ctx, menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func (s MenuService) Delete(ctx context.Context, id uint) error {
	return s.menus.Delete(ctx, id)
}

func (s MenuService) ListAll(ctx context.Context) ([]models.Menu, error) {
	return s.menus.ListAll(ctx)
}

func (s MenuService) ListByDate(ctx context.Context, date time.Time) ([]models.Menu, error) {
	return s.menus.ListByDate(ctx, date)
}

func buildMenu(input MenuInput) (*models.Menu, error) {
	title := strings.TrimSpace(input.Title)
	description := strings.TrimSpace(input.Description)
	if title == "" || description == "" || input.Price <= 0 || input.AvailableDate.IsZero() {
		return nil, ErrInvalidMenu
	}

	return &models.Menu{
		Title:         title,
		Description:   description,
		Price:         input.Price,
		AvailableDate: input.AvailableDate,
	}, nil
}
