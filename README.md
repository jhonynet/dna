# DNA Mutant Detector

Este algoritmo fue diseñado para evaluar una matriz NxN en todas las direcciones y detectar N secuencias de N caracteres (A, T, G, C) repetidos.

Para validar un ADN mutante, tiene que haber _**más de una secuencia de cuatro letras iguales, de forma oblicua, horizontal o vertical**_

Ya que solo se detecta secuencia de letras iguales y no palabras completas, se escanean solo 4 direcciones, evitando escanear dos veces lo mismo

## Constants
`const subSequencesToBeMutant int = 2`\
Cantidad de subsecuencias necesarias para que el ADN sea mutante

`const repeatCharsToBeMutant int = 4`\
Cantidad de caracteres repetidos para que una subsecuencia sea mutante

## Variables
`var mutantSubsequences int`\
Subsecuencias encontradas en la deteccion actual
 
`var directions = [4]Point{{1, 0}, {1, 1}, {0, 1}, {-1, 1}}`\
Todas las direcciones a escanear

## Types
Representacion de un punto en la matriz
```
type Point struct {
    x int
    y int
}
```

Representacion de un ADN\
```type Dna []string```


## Functions
* **func IsMutant(dna Dna) bool**\
Detecta si una ADN es mutante o humano

* **func BuildUniqueId(dna Dna) string**\
Genera un identificador unico para el ADN basado en sha1

* **func HasInvalidCharacters(dna Dna) bool**\
Detecta Caracteres invalidos en una matriz, actualmente los caracteres validos son A, T, G, C 

* **func IsSquareMatrix(dna Dna) bool**\
Detecta si la matriz introducida es de tipo NxN

## Testing & Bench
Este algoritmo tiene un coverage de 100% probando todas las funciones exportadas y privadas

| Test | n | ns/op |
| --- | --- | --- |
BenchmarkIsMutant-8|3000000|432 ns/op
BenchmarkIsSquareMatrix-8|300000000|5.22 ns/op
BenchmarkHasInvalidCharacters-8|30000000|55.4 ns/op
BenchmarkBuildUniqueId-8|3000000|577 ns/op

## Known Bugs
* Teniendo cantidad de repeticiones de caracteres configuradas por ejemplo en 4, cuando se detecta una subsecuencia de 5 caracteres iguales (en cualquier direccion), machea dos veces.
* En una matriz del doble de los caracteres requeridos para detectar subsecuencia mutante, habiendo macheado todos los caracteres (AAAAAAAA), se detecta una sola vez en lugar de 2