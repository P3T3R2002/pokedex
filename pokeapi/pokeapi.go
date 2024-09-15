package pokeapi

import (
	"errors"
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"github.com/P3T3R2002/pokedex/pokecache"
)

type Area struct{
	_		int 	`json:"count"`
	Next		string `json:"next"`
	Previous	string `json:"previous"`
	Results		[]Place `json:"results"`
}

type Place struct {
	Name	string `json:"name"`
	_		string `json:"url"`
}
     //\\
    //**\\
   //****\\
  //******\\

func Get_location(url string) Area {
	return Area{}
}

	//*********//

func Update_location(url string, area *Area) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//*********//
	if res.StatusCode >= 400 {
		return errors.New(fmt.Sprintf("error status: %s", res.Status))
	}
	//*********//
	data, err := io.ReadAll(res.Body)
	if err != nil {
			return err
	}
	//*********//
	err = json.Unmarshal(data, area)
	if err != nil {
		return err
	}
	//*********//
	return nil
}