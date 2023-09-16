open Recipientes
open RutaCorta

let pasosRecipientes = prof [2; 2] [2; 0]
printfn "%A" pasosRecipientes

let ruta = prof2 "i" "f" grafo
printfn "%A" ruta
0 // Salida del programa