service := Service{
    Name:   "api",
    Port:   8080,
    Status: "running",
}
servicePtr := &service
servicePtr.Status = "stopped"