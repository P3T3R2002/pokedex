package pokeapi

import (
	"errors"
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"github.com/P3T3R2002/pokedex/pokeapi/pokecache"
)

func Update_location(url string, area *pokecache.Area, cache *pokecache.Cache) error {
	ok ,err := get_from_cache(cache, url, area)
	if ok {
		return nil
	} else if err != nil {
		return err
	}

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
	add_to_cache(cache, url, area)
	return nil
}

//-----------------------------------------------------------------------
func Update_pockemon(url string, area *pokecache.Area, cache *pokecache.Cache) error {
	ok ,err := get_from_cache(cache, url, area)
	if ok {
		return nil
	} else if err != nil {
		return err
	}

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
	add_to_cache(cache, url, area)
	return nil
}

//-----------------------------------------------------------------------

func get_from_cache(cache *pokecache.Cache, url string, area *pokecache.Area) (bool, error) {
	err := pokecache.Read_cache(cache, url, area)
	if err != nil {
		if error.Error(err) == "Not in cache!" {
			return false, nil
		}
		return false, err
	} 
	return true, nil
}
//-----------------------------------------------------------------------

func add_to_cache(cache *pokecache.Cache, url string, area *pokecache.Area) {
	pokecache.Write_cache(cache, url, area)
	return 
}
