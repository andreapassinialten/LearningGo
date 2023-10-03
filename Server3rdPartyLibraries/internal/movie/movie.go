package movie

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`

	Cast []Actor `json:"cast,omitempty" gorm:"many2many:movie_cast;"` // Many-to-Many relationship
}

type Actor struct {
	ID         string `json:"id"`
	FirstName  string `json:"name"`
	SecondName string `json:"second_name"`
}

type Storage interface {
	CreateMovie(movie Movie) (Movie, error)
}

type Service struct {
	Storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		Storage: storage,
	}
}

func (s *Service) Create(m Movie) (Movie, error) {
	movie, err := s.Storage.CreateMovie(m)

	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
