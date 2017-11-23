/*

  Copyright (c) 2017, Vanessa Sochat
  All rights reserved.
  See LICENSE file in main repository

*/

package model

type db interface {
    SelectPeople() ([]*Person, error)
}
