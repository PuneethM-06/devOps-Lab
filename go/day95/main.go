func (s *Service)Stop(){
	s.Status = "stopped"
}

service.Stop()
fmt.Println(service.Status)