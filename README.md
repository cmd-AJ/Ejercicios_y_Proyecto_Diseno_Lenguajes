# Ejercicios_y_Proyecto_Diseno_Lenguajes


# Ejercicios_y_Proyecto_Diseno_Lenguajes

## Este proyecto tiene como objetivo la implementación de algoritmos para trabajar con autómatas finitos, conversion de expresiones regulares, incluyendo la generación de autómatas, su simulación, y la verificación de pertenencia de cadenas al lenguaje descrito por una expresión regular.


## ¿Como correr el Proyecto?


# Guía de Instalación

Para comenzar con este proyecto, sigue estos pasos para instalar las dependencias y configurar el entorno de desarrollo.

## Requisitos Previos

Antes de empezar, asegúrate de tener las siguientes herramientas instaladas:

- **Go** (Lenguaje de programación)
- **Nix** (Gestor de paquetes)

## Instalación de Dependencias

1. **Instalar Nix**  
   Sigue la [Guía de Instalación de Nix](https://nixos.org/download.html) para instalar Nix en tu sistema.

2. **Instalar Go**  
   Una vez que Nix esté instalado, usa Nix para descargar e instalar Go automáticamente:
   
   ```bash
   nix develop
## Construir y Ejecutar el Proyecto

Una vez que todas las dependencias estén instaladas, puedes proceder con la construcción y ejecución del proyecto.
Construir el Proyecto

Correr los comandos

  task build: Para construir el proyecto
  task run: Para correr el proyecto


## Categorias del Proyecto
*[Balanceador]([https://github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/tree/main/internal/balancer])
 - proporciona una implementación para determinar si una expresión regular está balanceada (contiene tanto los caracteres/simbolos de apertura como de cerradura).
*[Regex a Postfix]([https://github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/tree/main/internal/Postfix])
- Convierte a una cadena regex a un array o slice en caso de GO de simbolos a postfix con el algoritmo de Shunting Yard.
- Construye un AST a partir de una lista de símbolos en notación postfix.
*[AFD]([https://github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/tree/main/internal/dfa])
-Construccion de un automata a partir de una expresión Regex
*[AFD Minimizado]([https://github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/tree/main/internal/Minimal])
- Realiza el algoritmo de minimización de AFD, solo reduce los estados en el AFD si es que hay.




Resultados de una de las cadenas de prueba 

### Expresion regular: [a-d]ama\+

#### Resultados de prueba
![lab1sis](https://github.com/user-attachments/assets/27a147bb-bba2-46a0-8139-727df509f482)

Automata
![image](https://github.com/user-attachments/assets/9c771fc0-dd82-4971-bb09-b4e8be93bd12)

Atuomata Minimizado
![image](https://github.com/user-attachments/assets/fea3d976-523a-4610-8f4b-a3b287eba79c)



By Rayo and Jo
