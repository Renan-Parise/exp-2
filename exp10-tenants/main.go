package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID   int
	Name string
}

type Tenant struct {
	Name  string
	Users []User
}

type UserManager struct {
	Tenants map[string]Tenant
	mu      sync.Mutex
}

func NewUserManager() *UserManager {
	return &UserManager{
		Tenants: make(map[string]Tenant),
	}
}

func (um *UserManager) CreateTenant(name string) {
	um.mu.Lock()
	defer um.mu.Unlock()

	if _, ok := um.Tenants[name]; !ok {
		um.Tenants[name] = Tenant{Name: name}
		fmt.Printf("Tenant '%s' criado com sucesso.\n", name)
	} else {
		fmt.Printf("O Tenant '%s' já existe.\n", name)
	}
}

func (um *UserManager) AddUserToTenant(tenantName string, user User) {
	um.mu.Lock()
	defer um.mu.Unlock()

	if tenant, ok := um.Tenants[tenantName]; ok {
		tenant.Users = append(tenant.Users, user)
		um.Tenants[tenantName] = tenant
		fmt.Printf("Usuário '%s' adicionado ao tenant '%s'.\n", user.Name, tenantName)
	} else {
		fmt.Printf("O Tenant '%s' não existe.\n", tenantName)
	}
}

func (um *UserManager) GetUsersByTenant(tenantName string) {
	um.mu.Lock()
	defer um.mu.Unlock()

	if tenant, ok := um.Tenants[tenantName]; ok {
		fmt.Printf("Usuários no tenant '%s': %v\n", tenantName, tenant.Users)
	} else {
		fmt.Printf("O Tenant '%s' não existe.\n", tenantName)
	}
}

func main() {
	userManager := NewUserManager()

	userManager.CreateTenant("tA")
	userManager.CreateTenant("tB")

	userManager.AddUserToTenant("tA", User{ID: 1, Name: "u1"})
	userManager.AddUserToTenant("tA", User{ID: 2, Name: "u2"})
	userManager.AddUserToTenant("tB", User{ID: 3, Name: "u3"})

	userManager.GetUsersByTenant("tA")
	userManager.GetUsersByTenant("tB")
	userManager.GetUsersByTenant("tC")
}
