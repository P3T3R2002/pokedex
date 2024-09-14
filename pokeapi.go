package main

import (
	"errors"
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

type Area struct{
	Count		int 	`json:"count"`
	Next		string `json:"next"`
	Previous	string `json:"previous"`
	Results		[]Place `json:"results"`
}

type Place struct {
	Name	string `json:"name"`
	Url		string `json:"url"`
}
     //\\
    //**\\
   //****\\
  //******\\

func get_location(url string) (Area, error) {
	if url == "" {
		return Area{}, errors.New("No back")
	}
	res, err := http.Get(url)
	if err != nil {
		return Area{}, err
	}
	defer res.Body.Close()

	//*********//
	if res.StatusCode >= 400 {
		return Area{}, errors.New(fmt.Sprintf("error status: %s", res.Status))
	}
	//*********//
	data, err := io.ReadAll(res.Body)
	if err != nil {
			return Area{}, err
	}
	//*********//
	var area Area
	err = json.Unmarshal(data, &area)
	if err != nil {
		return Area{}, err
	}
	//*********//
	return area, nil
}

	//*********//

func update_location(url string, area *Area) error {
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