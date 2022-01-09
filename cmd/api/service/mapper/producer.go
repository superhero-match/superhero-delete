/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package mapper

import (
	am "github.com/superhero-match/superhero-delete/cmd/api/model"
	"github.com/superhero-match/superhero-delete/internal/producer/model"
)

// MapAPISuperheroToProducer maps API Superhero model to Producer Superhero model.
func MapAPISuperheroToProducer(superhero am.Superhero) model.Superhero {
	return model.Superhero{
		ID:        superhero.ID,
		DeletedAt: superhero.DeletedAt,
	}
}
