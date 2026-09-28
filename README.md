# Descripción del problema
Una esteticista solo ofrece la posibilidad de gestionear las citas presencialmente. Esto tiene una serie de inconvenientes durante su día a día:
  - **Interrupicón del servicio:** Tiene que interrumpir el servio que está realizando para gestionar las citas de otros clientes. 
  - **Pérdida de tiempo efectivo:** El tiempo utilizado en anotar nuevas citas se resta al tiempo asignado al tratamiento en curso. 
  - **Efecto dominó en la agenda:** La acumulación de retrasos desajusta la planificación del resto del día, afectando la puntulidad de las citas posteriores. 

## Caso concreto:
Es un martes a las 16:00. Eva, la esteticista, está realizando un servicio de depilación de cejas con una duración estimada de 15 minutos (finalización prevista a las 16:15).

A las 16:10, llegan dos clientas al local (Rosa y Ana) para pedir cita. Eva se ve obligada a pausar el servicio en curso para agendar el servicio de uñas de Rosa y la sesión de láser de Ana.

Cuando Eva reanuda la depilación de cejas son las 16:15, hora en la que el servicio debería haber concluido. Como todavía le restan 7 minutos de trabajo, finaliza a las 16:22. Dado que la siguiente cita estaba programada para las 16:20, el horario del resto de la tarde queda desfasado, acumulando retrasos y tiempos de espera para las clientas posteriores.

![Fotografía de la tarjeta de rol](imagenes/cliente.jpeg)

# Como solucionarlo como desarrollador
Para resolver este problema, la propuesta consiste en desarrollar una plataforma web o aplicación móvil que permita a los clientes consultar los servicios disponibles, seleccionar una fecha y hora, y reservar de forma autónoma.

## Lógica de negocio
  - La selección de días y horas debe restringirse al horario y jornada laboral configurado por la esteticista.
  - El sistema debe garantizar que no se reserven dos servicios en un mismo intervalo de tiempo.
  - La asignación de citas debe implementar un algoritmo o conjunto de reglas que reduzca los tiempos ociosos entre servicios consecutivos.

# Datos necesarios para la solución
Los datos necesarios para poder solucionar el problema, tales como el horario laboral o los servicios ofertados con su respectivo precio y duración, se obtendrán directamente de la esteticista, por tanto no será necesario obtener ninguna información de fuentes externas.

Para el almacenamiento y gestión de los servicios se utilizará una hoja de cálculo. En ella, cada servicio tendrá asociado su precio y duración estimada, permitiendo colsultar y extraer los datos necesarios en cualquier momento.

# Servicios ofertados
## Servicios faciales
  - Limpieza facial 
  - Tratamiento facial
## Maquillajes
  - Maquillaje de día
  - Maquillade de madrina
  - Maquillaje de novia
## Depilación con cera
  - Labios y cejas
  - Medias piernas
  - Depilación completa
## Depilación láser
  - Ingles y axilas
  - Piernas enteras
  - Espalda

Para consultar la duración estimada y precio de cada servicio revisar [servicios_ofertados.xlsx](./servicios_ofertados.xlsx). El tiempo estimado es vital para la lógica de negocio, para evitar solapamiento de servicios. El precio es dato de alto interés para los usuarios. 

# Configuración inicial
Toda la configuración inicial se encuentra en [Configuracion.md](./Configuracion.md)
