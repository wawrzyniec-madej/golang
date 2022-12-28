package finders

import "gin-tutorial/dtos"

func FindAlbumByIdInAlbums(albumID string) *dtos.AlbumDTO {

	for _, albumDTO := range dtos.Albums {
		if albumDTO.ID == albumID {
			return &albumDTO
		}
	}

	return nil
}
