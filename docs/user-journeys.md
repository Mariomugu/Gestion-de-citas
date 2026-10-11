# User journeys
## 1: Planificar la semana siguiente

 - **Frecuencia:** Una vez por semana.
 - **Contexto:** En casa o fuera del trabajo.
 - **Dispositivo:** Portátil o móvil (el que prefiera en el momento).

Cada lunes por la tarde, Eva, la esteticista, planificará la próxima semana con las citas que haya ido acumulando hasta dicho día. Eva introducirá las citas solicitadas indicando el cliente, los días u horas de disponibilidad del cliente y el servicio solicitado. El sistema devolverá la planificicación óptima
que aprovecha mejor el tiempo minimizando los tiempos de preparación y limpieza. Eva una vez que disponga de la planificación de dicha semana puede comunicar a cada cliente su día y hora concreta.
 
## 2: Añadir una cita durante la semana en un hueco libre

 - **Frecuencia** Varias veces por semana.
 - **Contexto** Una clienta pide cita después de publicado el plan. La solicitud llega en cualquier momento y Eva puede estar en casa, en la calle o atendiendo a otra clienta.
 - **Dispositivo** Móvil.

Eva abre el sistema e indica el servicio solicitado y la disponibilidad de la clienta. El sistema busca los huecos en los que la cita quepa entera. El sistema le muestra las opciones ordenadas por menos tiempo perdido. Por ejemplo, para una cita de cera, un hueco justo después de otra cita de cera cuesta 0 minutos extra de preparación, mientras que tras un tratamiento facial implicaría 20 minutos extra de preparación. Eva elige una opción y el sistema añade la cita al plan sin mover ninguna cita ya confirmada con anterioridad. Finalmente, Eva puede responder a esta nueva clienta con el día y hora de su cita.

### Caso alternativo: no hay hueco
El sistema indica que no hay hueco disponible, y añade dicha cita a una lista de espera. En caso de que haya alguna cancelación se podrá llenar el hueco con las citas en dicha lista. Eva comunica a la clienta que ha entrado en dicha lista y también le ofrece la posibilidad de añadir nueva disponibilidad para dar una cita en la planificación de semanas posteriores.
 
## 3: Cancelación y aprovechamiento del hueco

 - **Frecuencia** Ocasional.
 - **Contexto** Una clienta cancela.
 - **Dispositivo** Móvil.

Eva accede al sistema y marca la cita como cancelada. El sistema libera el hueco y muestra cuántos minutos quedan libres en ese turno y si se han fusionado varios huecos. Propone clientas de la lista de espera que caben en ese hueco y cuáles encajan mejor por maquinaria. Eva elige una(o varias si el hueco lo permite) y la cita se añade, o decide dejar el hueco libre. Si Eva ha decidido elegir alguna cita se lo comunica al cliente correspondiente para saber si aún está disponible para dicha cita. En caso contrario, volvería a marcar esta nueva cita como cancelada y repetiría el proceso.
