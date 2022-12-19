create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('A Virgem da Conceição', 'A Virgem da Conceição, também identificada como Rainha da Floresta, tornou-se patrona da religião do Santo Daime em função de sua presença nas narrativas fundacionais que a identificam como instrutora espiritual do fundador Raimundo Irineu Serra (Mestre Irineu) em sua iniciação com a ayahuasca, bebida de origem indígena e utilizada no culto daimista.'),
('Mestre Irineu', 'Raimundo Irineu Serra conheceu por meio de indígenas da região a bebida sagrada ayahuasca.'),
('Padrinho Sebastião', 'Sebastião Mota de Melo foi discípulo direto de Raimundo Irineu Serra, a cujo trabalho deu continuidade em uma doutrina que usa como sacramento a bebida chamada ayahuasca, batizada por Raimundo de Daime, associada a orações e cânticos (hinos) a diversas divindades.'),
('Mestre Gabriel', 'José Gabriel da Costa foi seringueiro, “soldado da borracha”. Em 1959, quando trabalhava nos seringais do Acre, na fronteira entre o Brasil e a Bolívia, entrou em contato com o chá Hoasca. Lá, começou a distribuir o chá, formando os seus primeiros discípulos. Em 22 de julho de 1961, ainda dentro da floresta, declarou a criação da União do Vegetal, religião que tem como símbolo a Luz, a Paz e o Amor.'),
('Mestre Francisco', 'Francisco Souza de Almeida foi o fundador do Centro de Cultura Cósmica. Nascido no Estado do Ceará, Mestre Francisco viveu no Estado do Acre, onde conheceu o chá Ayahuasca e conviveu com o Mestre Raimundo Irineu Serra, fundador da linha do Daime e precursor do uso do chá entre as doutrinas religiosas. Também conviveu com grandes líderes ayahuasqueiros como o Padrinho Sebastião, do Santo Daime, o Mestre Antônio Geraldo, da Barquinha, e acompanhou o Mestre Gabriel, fundador da União do Vegetal, tendo sido filiado neste centro.');