-----------------------------------------------------------------------
Consideraciones generales
-----------------------------------------------------------------------
Muchas gracias por la oportunidad.

Funcionalidades:
	Alta de sucursales
	Consulta de una sucursal
	Consulta de todas las sucursales
	Consulta de sucursal mas cercana
	
Base de datos:
	Elegí usar SQL Lite ya que puede ser embebida en la aplicación de modo que pueda ser probada tal y como está, sin configurar y crear una base y la tabla.
	
-----------------------------------------------------------------------
Consideraciones para la instalación
-----------------------------------------------------------------------
Instalar éstos paquetes:

	go get github.com/mattn/go-sqlite3
	
	go get github.com/gorilla/mux
	
-----------------------------------------------------------------------
Enpoints
-----------------------------------------------------------------------
GETs:

	http://localhost:10000/
		Página home con texto.
		
	http://localhost:10000/all
		Muestra todas las sucursales existetes en la base de datos.
		
	http://localhost:10000/sucursal/{id}
		Muestra la sucursal correspondiente al id que se requiere. Ejemplo:
			http://localhost:10000/sucursal/1
POSTs:

	http://localhost:10000/altasucursal
		Graba una sucursal en la base de datos. Ejemplo:
			{
				"Id": "1", 
				"Direccion": "Martinez", 
				"Latitud": "-34.511120", 
				"Longitud": "-58.503884" 
			}
			
	http://localhost:10000/sucursalmascercana
		Recibe coordeadas de latitud y longitud y muestra la sucursal mas cercana. Ejemplo:
			{
				"Lat": "-14.5684", 
				"Long": "175.472636"
			}

-----------------------------------------------------------------------
Datos de prueba
-----------------------------------------------------------------------
Sucursales de prueba para cargar en la base:

	{
		"Id": "1", 
		"Direccion": "Martinez", 
		"Latitud": "-34.511120", 
		"Longitud": "-58.503884" 
	}


	{
		"Id": "2", 
		"Direccion": "Microcentro", 
		"Latitud": "-34.603033", 
		"Longitud": "-58.376234" 
	}

	{
		"Id": "3", 
		"Direccion": "La Plata", 
		"Latitud": "-34.926463", 
		"Longitud": "-57.969472" 
	}

	{
		"Id": "4", 
		"Direccion": "Tierra del fuego", 
		"Latitud": "-53.745054", 
		"Longitud": "-67.767361"
	}

	{
		"Id": "5", 
		"Direccion": "Jujuy", 
		"Latitud": "-23.536447", 
		"Longitud": "-65.381489" 
	}
	
-----------------------------------------------------------------------
Puntos de prueba para enviar POST a la aplicación en http://localhost:10000/sucursalmascercana:
-----------------------------------------------------------------------

Caso exitoso:

	{
		"Lat": "-34.602126", 
		"Long": "-58.393904"
	}
	
Caso NO exitoso:

	{
		"Lat": "aaaaaa", 
		"Long": "175.472636"
	}
