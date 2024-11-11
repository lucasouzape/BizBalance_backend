package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model          // inclui campos como ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Modelo para a tabela administrador
type Administrador struct {
	gorm.Model
	Nome  string `json:"nome" gorm:"size:100;not null"`
	Email string `json:"email" gorm:"size:100;unique;not null"`
	Senha string `json:"senha" gorm:"size:100;not null"`
}

// Modelo para a tabela pao_de_mel
type PaoDeMel struct {
	gorm.Model
	Sabor           string        `json:"sabor" gorm:"size:50;not null"`
	Quantidade      int           `json:"quantidade" gorm:"not null"`
	Validade        time.Time     `json:"validade" gorm:"not null"`
	PrecoCusto      float64       `json:"preco_custo" gorm:"type:numeric(10,2);not null"`
	PrecoVenda      float64       `json:"preco_venda" gorm:"type:numeric(10,2);not null"`
	AdministradorID uint          `json:"administrador_id"`
	Administrador   Administrador `json:"administrador" gorm:"foreignKey:AdministradorID"`
}

// Modelo para a tabela pao_de_mel_2
type PaoDeMel2 struct {
	gorm.Model
	PaoDeMelID      uint          `json:"pao_de_mel_id" gorm:"not null"`
	PaoDeMel        PaoDeMel      `json:"pao_de_mel" gorm:"foreignKey:PaoDeMelID"`
	Sabor           string        `json:"sabor" gorm:"size:50;not null"`
	Quantidade      int           `json:"quantidade" gorm:"not null"`
	Validade        time.Time     `json:"validade" gorm:"not null"`
	PrecoCusto      float64       `json:"preco_custo" gorm:"type:numeric(10,2);not null"`
	PrecoVenda      float64       `json:"preco_venda" gorm:"type:numeric(10,2);not null"`
	AdministradorID uint          `json:"administrador_id"`
	Administrador   Administrador `json:"administrador" gorm:"foreignKey:AdministradorID"`
}

// Modelo para a tabela brownie
type Brownie struct {
	gorm.Model
	Sabor           string        `json:"sabor" gorm:"size:50;not null"`
	Quantidade      int           `json:"quantidade" gorm:"not null"`
	Validade        time.Time     `json:"validade" gorm:"not null"`
	PrecoCusto      float64       `json:"preco_custo" gorm:"type:numeric(10,2);not null"`
	PrecoVenda      float64       `json:"preco_venda" gorm:"type:numeric(10,2);not null"`
	AdministradorID uint          `json:"administrador_id"`
	Administrador   Administrador `json:"administrador" gorm:"foreignKey:AdministradorID"`
}

// Modelo para a tabela recheio
type Recheio struct {
	gorm.Model
	PaoDeMelID      uint          `json:"pao_de_mel_id" gorm:"not null"`
	PaoDeMel        PaoDeMel      `json:"pao_de_mel" gorm:"foreignKey:PaoDeMelID"`
	Sabor           string        `json:"sabor" gorm:"size:50;not null"`
	Quantidade      int           `json:"quantidade" gorm:"not null"`
	Validade        time.Time     `json:"validade" gorm:"not null"`
	PrecoCusto      float64       `json:"preco_custo" gorm:"type:numeric(10,2);not null"`
	PrecoVenda      float64       `json:"preco_venda" gorm:"type:numeric(10,2);not null"`
	AdministradorID uint          `json:"administrador_id"`
	Administrador   Administrador `json:"administrador" gorm:"foreignKey:AdministradorID"`
}
