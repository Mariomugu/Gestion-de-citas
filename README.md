# Riesgo vial urbano para ciclistas y usuarios de patinete según franja horaria y clima
# Descripción del problema
Las personas que se desplazan a diario en bicicleta o en patinete (VMP) por una ciudad eligen su ruta con criterios como la duración o la comodidad, pero no saben en qué calles, a qué horas y con qué clima se acumulan los accidentes que afectan a su tipo de vehículo. Como no tienen ese dato, toman decisiones a ciegas: una calle que parece tranquila puede concentrar muchos accidentes en la franja en la que la usan, y otra que les da miedo puede ser en la práctica menos problemática.

## Caso concreto:
Es un lunes de octubre por la mañana y está lloviendo. Lucía va en bici al trabajo y tiene tres formas de llegar, que pasan por la calle A, la calle B o la calle C. Sale a las 8:15, en hora punta. No sabe si alguna de esas calles tiene un historial de accidentes con bicicletas en esa franja y con lluvia, ni si es peor que las otras dos. Hoy elegirá por intuición, y quizá acabe usando la calle donde más ciclistas han sufrido accidentes en días como este.

![Fotografía de la tarjeta de rol](imagenes/cliente.jpeg)

# Lógica de negocio
  - **Extraer** los registros del fichero CSV.
  - **Analizar** cada registro para agrupar las filas que comparten número de expediente y obtener así accidentes distintos, ya que cada fila es una persona implicada y un mismo accidente aparece varias veces.
  - **Filtrar** los accidentes en los que al menos una persona implicada iba en bicicleta o en VMP, según el tipo de vehículo de cada fila.
  - **Validar** los registros, descartando o marcando como desconocidos los que carecen de hora, de estado meteorológico o de localización.
  - **Normalizar** los nombres de calle del fichero y los escritos por el usuario (mayúsculas, tildes, abreviaturas como "CALL.", "AVDA." o "GTA.") para que puedan compararse.
  - **Calcular** un coeficiente de riesgo por calle, franja horaria y estado meteorológico, a partir del número de accidentes en esa combinación.
  - **Generar** una calificación de peligrosidad para cada calle de la lista, comparando su coeficiente con el del conjunto de calles de la ciudad, e indicando "sin datos suficientes" cuando la calle apenas tenga registros.

# Datos necesarios para la solución
Se van a utilizar los datos de accidentes de tráfico de la ciudad de Madrid disponibles en el [Portal de datos abiertos del Ayuntamiento de Madrid](https://datos.madrid.es/dataset/300228-0-accidentes-trafico-detalle/downloads?hierarchy=2019), se utilizarán los archivos en formato csv del año 2019 en adelante.

## Ejemplo de entrada en los ficheros
```csv
num_expediente;fecha;hora;localizacion;numero;cod_distrito;distrito;tipo_accidente;estado_meteorológico;tipo_vehiculo;tipo_persona;rango_edad;sexo;cod_lesividad;lesividad;coordenada_x_utm;coordenada_y_utm;positiva_alcohol;positiva_droga
2018S017842;04/02/2019;9:10:00;CALL. ALBERTO AGUILERA, 1;1;1;CENTRO;Colisión lateral;Despejado;Motocicleta > 125cc;Conductor;De 45 a 49 años;Hombre;7;Asistencia sanitaria sólo en el lugar del accidente;440068;4475679;N;NULL
```

# Configuración inicial
Toda la configuración inicial se encuentra en [Configuracion.md](./Configuracion.md)
