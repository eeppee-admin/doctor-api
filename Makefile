all:
	go run .\cmd\

note:
	@echo [GIN-debug] POST   /api/users/register      
	@echo [GIN-debug] POST   /api/users/login          
	@echo [GIN-debug] GET    /api/patients/:id         
	@echo [GIN-debug] GET    /api/patients/:id/records 
	@echo [GIN-debug] GET    /api/medicines           
	@echo [GIN-debug] GET    /api/medicines/:id        
	@echo [GIN-debug] POST   /api/medicines            
	@echo [GIN-debug] PUT    /api/medicines/:id        
	@echo [GIN-debug] DELETE /api/medicines/:id        