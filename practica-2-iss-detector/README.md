# Iss Detector

La Estación Espacial Internacional (ISS) orbita a una velocidad de 27.580 km/h. Tarda 90 minutos en dar una vuelta 
alrededor de la tierra.

Existen dos APIs públicas que nos puede dar información acerca de su posición en cada instante 
http://api.open-notify.org/iss-now.json y https://wheretheiss.at/w/developer. Con la longitud y latitud 
podemos determinar el país que está sobrevolando utilizando otra API: http://ws.geonames.org.

El obejtivo de esta práctica es crear un programa que nos informe cada vez que la ISS cambia de país haciendo una 
comprobación cada 10 segundos. 

El output será del estilo: 

```bash
> go run main.go --tracker [iss-now|whereisiss]
Afganistan
China
Corea del Norte
Japón
....
```

Con un parámetro podemos decidir si utilizar iss-now o whereisiss, así que es necesario implementar 
el patrón Strategy e implementar una interfaz común a los dos trackers.



