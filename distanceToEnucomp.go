package distanceToEnucomp

import "math"

func Info() string {
	return "O ENUCOMP vem se firmando como um dos maiores eventos de tecnologia do estado do Piauí e " +
		" estados vizinhos da região Nordeste. A edição de 10 anos será realizada de 15 a 17 de novembro na " +
		"cidade de Parnaíba-PI, porta de entrada turística para a Rota das Emoções (Lençóis Maranhenses - Delta" +
		" do Rio Parnaíba - Jericoacoara). Sua programação contará com palestrantes nacionais, minicursos, " +
		"sessões técnicas de apresentação de artigos científicos (oral e painel), keynotes e a exposição de aplicativos."
}

func ShortCourses() []string {
	s := []string{
		"Minicurso I - Softwarização em Redes: Do Plano de Dados ao Plano de Orquestração",
		"Minicurso II - Roteamento Dinâmico em IoT baseado em Requisitos de Aplicações Específicas",
		"Minicurso III - Construção de Aplicações baseadas em Física: Uma introdução ao ODE",
		"Minicurso IV - Processamento Digital de Imagens Médicas com Python e Opencv",
		"Minicurso V - Produção de Objetos de Aprendizagem Baseados em Vídeos Interativos voltados para o ambiente TVDi",
		"Minicurso VI - Uma Introdução ao Go: A Linguagem Performática do Google",
		"Minicurso VII - Uma Introdução à Robótica Móvel",
		"Minicurso VIII - Construindo Microsserviços em Ambiente de Computação em Nuvem",
		"Minicurso IX - Deep Learning: Uma Introdução às Redes Neurais Convolucionais",
		"Minicurso X - Introdução ao Desenvolvimento de Aplicativos Android Utilizando Conceitos de Geolocalização",
		"Minicurso XI - Monitoramento de Tráfego em Redes de Internet das Coisas",
		"Minicurso XII - Introdução a Análise de Dados com Python e Pandas",
	}
	return s
}

// -2.909433, -41.754527 UFPI-PHB
func DistanceTo(lat, lon float64) float64 {
	R := 6371.0
	latUFPIPHB := -2.909433
	lonUFPIPHB := -41.754527
	dLat := (latUFPIPHB - lat) * (math.Pi / 180)
	dLon := (lonUFPIPHB - lon) * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat*(math.Pi/180))*
		math.Cos(-7.081752*(math.Pi/180))*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := R * c
	return d
}
