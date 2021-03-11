package repositories

import (
	"log"
	. "star-wars-api/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PlanetRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "planets"
)

func (p *PlanetRepository) Connect() {
	session, err := mgo.Dial(p.Server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(p.Database)
}

func (p *PlanetRepository) GetAll() ([]Planet, error) {
	var planets []Planet
	err := db.C(COLLECTION).Find(bson.M{}).All(&planets)

	return planets, err
}

func (p *PlanetRepository) GetById(id string) (Planet, error) {
	var planet Planet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&planet)

	return planet, err
}

func (p *PlanetRepository) GetByName(name string) (Planet, error) {
	var planet Planet
	err := db.C(COLLECTION).Find(bson.M{"name": name}).One(&planet)

	return planet, err
}

func (p *PlanetRepository) Create(planet Planet) error {
	err := db.C(COLLECTION).Insert(&planet)

	return err
}

func (p *PlanetRepository) Update(id string, planet Planet) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &planet)

	return err
}

func (p *PlanetRepository) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))

	return err
}
