 #  Hackers E-commerce

Este proyecto consiste en el diseño y desarrollo del sistema de gestión para una tienda en línea de accesorios de computadora llamada **"Hackers"**. Desarrollado como parte del plan de estudios de la UIDE (Modalidad a Distancia), el sistema aplica los principios fundamentales de la **Programación Funcional** utilizando el lenguaje Go.

---

##  Arquitectura del Proyecto

El backend está estructurado de forma modular para garantizar la escalabilidad y un bajo acoplamiento entre sus componentes:

- `main.go`: Punto de entrada principal que inicializa y arranca el servidor del sistema.
- `paquetes/catalogo/`: Módulo encargado de la gestión inmutable del inventario de accesorios de hardware.
- `paquetes/carrito/`: Módulo enfocado en operaciones puras para la manipulación de la cesta de compras.
- `paquetes/checkout/`: Módulo de facturación y cálculo dinámico de impuestos/descuentos mediante funciones de orden superior.

---

##  Enfoque de Programación Funcional

A diferencia del paradigma imperativo tradicional, este backend implementa:
1. **Funciones Puras:** Operaciones que no producen efectos secundarios y devuelven siempre el mismo resultado para los mismos argumentos.
2. **Inmutabilidad:** Las colecciones de datos (como el inventario o el carrito) no se modifican directamente; en su lugar, se generan nuevas instancias con los cambios reflejados.
3. **Funciones de Orden Superior & Closures:** Utilizadas principalmente en el módulo de checkout para crear reglas de facturación dinámicas y aplicar descuentos personalizados de manera declarativa.

---

##  Instrucciones de Ejecución

Para levantar y probar el servidor localmente, ejecuta el siguiente comando en la raíz del proyecto:

```bash
go run "package main.go"
