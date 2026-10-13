# Historias de usuario
## HU001: No sé cuál es el mejor orden para las citas
Eva es esteticista autónoma y trabaja sola. Ofrece varios servicios en un mismo espacio y cada tipo de servicio usa un equipamiento distinto (facial, maquillaje, láser o cera). Cada equipamiento necesita un tiempo de preparación antes de usarlo y un tiempo de limpieza después. Cuando entre dos citas seguidas cambia de tipo de servicio, tiene que recoger el equipamiento anterior y preparar el siguiente desaprovechando tiempo útil.

Hoy asigna las citas por orden de llegada, porque no tiene una forma sencilla de saber qué orden le haría perder menos tiempo en esos cambios. Además tiene que hacer dicha planificación óptima con una semana de antelación, para que cada clienta conozca su hora concreta con margen.

Datos necesarios: [horario](../datos/horario.csv), [maquinaria](../datos/maquinaria.csv) y [servicios](../datos/servicios.csv)

[Jornada 1](./user-journeys.md)

## HU002: No sé si puedo añadir nuevas citas en la semana actual ya planificada
Una vez confirmado el plan de la semana, Eva sigue recibiendo solicitudes de clientas. Algunas encajarán en los huecos que quedaron libres(si hay). Otras pueden
encajar porque una clienta cancela y deja un hueco que antes no existía. En ambos casos, Eva deber ser capaz de dar una hora a la nueva clienta sin cambiar la hora de las que ya la conocen y manteniendo bajos los tiempos de preparación y recogida. Hoy no sabe con rapidez si una cita nueva cabe, ni cuál es el mejor hueco para ella.

[Jornadas 2 y 3](./user-journeys.md)