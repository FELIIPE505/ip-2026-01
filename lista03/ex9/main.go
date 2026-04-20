package main

import f "fmt"

func main() {
	var n1, n2, somaalunos, media float64
	var N, aprovado, reprovado, exame int
	var mediaaluno []float64
	somaalunos = 0

	f.Println("Insira a quantidade de alunos:")
	f.Scan(&N)

	for i := 0; i < N; i++ {
		f.Printf("Insira a nota 1 do aluno %d: ", i+1)
		f.Scan(&n1)
		f.Printf("Insira a nota 2 do aluno %d: ", i+1)
		f.Scan(&n2)
		if n1 >= 0 && n1 <= 10 && n2 >= 0 && n2 <= 10 {

			media = (n1 + n2) / 2
			mediaaluno = append(mediaaluno, media)
			somaalunos = somaalunos + mediaaluno[i]
			if mediaaluno[i] >= 0 && mediaaluno[i] < 3 {
				reprovado = reprovado + 1
				f.Println("REPROVADO")
			}
			if mediaaluno[i] >= 3.0 && mediaaluno[i] <= 7.0 {
				exame = exame + 1
				f.Println("EXAME")
			}
			if mediaaluno[i] > 7 {
				aprovado = aprovado + 1
				f.Println("APROVADO")
			}
		} else {
			f.Println("Nota inserida inválida")
			return
		}
	}
	f.Println("O total de alunos aprovados é:", aprovado)
	f.Println("O total de alunos reprovados é:", reprovado)
	f.Println("O total de alunos de exames é:", exame)
	f.Println("A média da classe é:", somaalunos/float64(N))
}
