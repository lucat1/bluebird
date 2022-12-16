package request

var allPoliticians = [MAXPOLITICIANS][2]string{{"Andrea", "Abodi"}, {"Maurizio", "Acerbo"}, {"Francesco", "Acquaroli"}, {"Mario", "Adinolfi"}, {"Davide", "Aiello"}, {"Lucia", "Albano"}, {"Alessandro", "Alfieri"}, {"Enrica", "Alifano"}, {"Cristina", "Almici"}, {"Vincenza", "Aloisio"}, {"Gaetano", "Amato"}, {"Paola", "Ambrogio"}, {"Alessia", "Ambrosi"}, {"Vincenzo", "Amendola"}, {"Enzo", "Amich"}, {"Bartolomeo", "Amidei"}, {"Alessandro", "Amorese"}, {"Renato", "Ancorotti"}, {"Giorgia", "Andreuzza"}, {"Antonio", "Angelucci"}, {"Alfredo", "Antoniozzi"}, {"Chiara", "Appendino"}, {"Giovanni", "Arruzzolo"}, {"Anna", "Ascani"}, {"Stefania", "Ascari"}, {"Bruno", "Astorre"}, {"Andrea", "Augello"}, {"Carmela", "Auriemma"}, {"Alberto", "Bagnai"}, {"Roberto", "Bagnasco"}, {"Ouiddad", "Bakkali"}, {"Alberto", "Balboni"}, {"Antonio", "Baldelli"}, {"Vittoria", "Baldino"}, {"Andrea", "Barabotti"}, {"Alberto", "Barachini"}, {"Anthony", "Barbagallo"}, {"Michele", "Barcaiuolo"}, {"Vito", "Bardi"}, {"Paolo", "Barelli"}, {"Valentina", "Barzotti"}, {"Lorenzo", "Basso"}, {"Alessandro", "Battilocchio"}, {"Francesco", "Battistoni"}, {"Alfredo", "Bazoli"}, {"Davide", "Bellomo"}, {"Maria", "Bellucci"}, {"Stefano", "Benigni"}, {"Domenico", "Bennardi"}, {"Gostoli", "Benvenuti"}, {"Alessandro", "Benvenuto"}, {"Fabrizio", "Benzoni"}, {"Davide", "Bergamini"}, {"Deborah", "Bergamini"}, {"Giorgio", "Bergesio"}, {"Silvio", "Berlusconi"}, {"Anna", "Bernini"}, {"Giovanni", "Berrino"}, {"Mauro", "Berruto"}, {"Pier", "Bersani"}, {"Dolores", "Bevilacqua"}, {"Michaela", "Biancofiore"}, {"Giuseppe", "Bicchielli"}, {"Matteo", "Biffoni"}, {"Galeazzo", "Bignami"}, {"Simone", "Billi"}, {"Anna", "Bilotti"}, {"Rosy", "Bindi"}, {"Pierluigi", "Biondi"}, {"Ingrid", "Bisa"}, {"Massimo", "Bitonci"}, {"Mara", "Bizzotto"}, {"Francesco", "Boccia"}, {"Gianangelo", "Bof"}, {"Laura", "Boldrini"}, {"Stefano", "Bonaccini"}, {"Simona", "Bonafè"}, {"Angelo", "Bonelli"}, {"Elena", "Bonetti"}, {"Giulia", "Bongiorno"}, {"Francesco", "Bonifazi"}, {"Emma", "Bonino"}, {"Simona", "Bordonali"}, {"Mario", "Borghese"}, {"Stefano", "Borghesi"}, {"Claudio", "Borghi"}, {"Enrico", "Borghi"}, {"Lucia", "Borgonzoni"}, {"Francesco", "Borrelli"}, {"Maria", "Boschi"}, {"Umberto", "Bossi"}, {"Amedeo", "Bottaro"}, {"Chiara", "Braga"}, {"Michela", "Brambilla"}, {"Beatrice", "Brignone"}, {"Luigi", "Brugnaro"}, {"Giovanna", "Bruno"}, {"Raffaele", "Bruno"}, {"Francesco", "Bruzzone"}, {"Carmela", "Bucalo"}, {"Marco", "Bucci"}, {"Alice", "Buonguerrieri"}, {"Alessio", "Butti"}, {"De", "Cafiero"}, {"Salvatore", "Caiata"}, {"Nicola", "Calandrini"}, {"Paolo", "Calcinaro"}, {"Roberto", "Calderoli"}, {"Marina", "Calderone"}, {"Tommaso", "Calderone"}, {"Carlo", "Calenda"}, {"Giangiacomo", "Calovini"}, {"Susanna", "Campione"}, {"Nanni", "Campus"}, {"Susanna", "Camusso"}, {"Stefano", "Candiani"}, {"Alessandro", "Canelli"}, {"Gerolamo", "Cangiano"}, {"Giovanni", "Cannata"}, {"Francesco", "Cannizzaro"}, {"Gianluca", "Cantalamessa"}, {"Luciano", "Cantone"}, {"Maria", "Cantù"}, {"Virginio", "Caparvi"}, {"Ugo", "Cappellacci"}, {"Enrico", "Cappelletti"}, {"Renzo", "Caramaschi"}, {"Alessandro", "Caramiello"}, {"Maria", "Caretta"}, {"Mara", "Carfagna"}, {"Mirco", "Carloni"}, {"Ida", "Carmina"}, {"Andrea", "Caroppo"}, {"Dario", "Carotenuto"}, {"Franz", "Caruso"}, {"Maurizio", "Casasco"}, {"Maria", "Casellati"}, {"Pier", "Casini"}, {"Antonio", "Caso"}, {"Giuseppe", "Cassì"}, {"Guido", "Castelli"}, {"Maria", "Castellone"}, {"Francesco", "Castiello"}, {"Enrico", "Castiglione"}, {"Piero", "Castrataro"}, {"Andrea", "Casu"}, {"Roberto", "Cataldi"}, {"Alessandro", "Cattaneo"}, {"Elena", "Cattaneo"}, {"Vanessa", "Cattoi"}, {"Laura", "Cavandoli"}, {"Ilaria", "Cavo"}, {"Fabrizio", "Cecchetti"}, {"Giulio", "Centemero"}, {"Gian", "Centinaio"}, {"Marco", "Cerreto"}, {"Lorenzo", "Cesa"}, {"Susanna", "Cherchi"}, {"Paola", "Chiesa"}, {"Monica", "Ciaburro"}, {"Francesco", "Ciancitto"}, {"Paolo", "Ciani"}, {"Luciano", "Ciocchetti"}, {"Alessandro", "Ciriani"}, {"Luca", "Ciriani"}, {"Edmondo", "Cirielli"}, {"Pomicino", "Cirino"}, {"Alberto", "Cirio"}, {"Dimitri", "Coin"}, {"Mariolina", "Colangelo"}, {"Marta", "Collot"}, {"Beatriz", "Colombo"}, {"Chiara", "Colosimo"}, {"Alessandro", "Colucci"}, {"Alfonso", "Colucci"}, {"Silvana", "Comaroli"}, {"Fabrizio", "Comba"}, {"Saverio", "Congedo"}, {"Giuseppe", "Conte"}, {"Mario", "Conte"}, {"Michele", "Conti"}, {"Marcello", "Coppo"}, {"Claudio", "Corradino"}, {"Andrea", "Corsaro"}, {"Piergiorgio", "Cortelazzo"}, {"Giulia", "Cosenza"}, {"Enrico", "Costa"}, {"Sergio", "Costa"}, {"Carlo", "Cottarelli"}, {"Stefania", "Craxi"}, {"Andrea", "Crippa"}, {"Andrea", "Crisanti"}, {"Marco", "Croatti"}, {"Guido", "Crosetto"}, {"Ilaria", "Cucchi"}, {"Gianni", "Cuperlo"}, {"Augusto", "Curti"}, {"Di", "Cuttica"}, {"Gianguido", "D'Alberto"}, {"Massimo", "D'Alema"}, {"Antonio", "D'Alessio"}, {"Luciano", "D'Alfonso"}, {"Mauro", "D'Attis"}, {"Cecilia", "D'Elia"}, {"Valentina", "D'Orso"}, {"Chiesa", "Dalla"}, {"Concetta", "Damante"}, {"Dario", "Damiani"}, {"Andrea", "Dara"}, {"Bertoldi", "De"}, {"Carlo", "De"}, {"Corato", "De"}, {"Cristofaro", "De"}, {"Luca", "De"}, {"Luca", "De"}, {"Magistris", "De"}, {"Maria", "De"}, {"Micheli", "De"}, {"Monte", "De"}, {"Mossi", "De"}, {"Palma", "De"}, {"Pascale", "De"}, {"Pellegrin", "De"}, {"Poli", "De"}, {"Priamo", "De"}, {"Rosa", "De"}, {"Antonio", "Decaro"}, {"Salvatore", "Deidda"}, {"Barba", "Del"}, {"Bono", "Del"}, {"Porta", "Della"}, {"Vedova", "Della"}, {"Delle", "Delmastro"}, {"Graziano", "Delrio"}, {"Battista", "Di"}, {"Biase", "Di"}, {"Girolamo", "Di"}, {"Giusepe", "Di"}, {"Lauro", "Di"}, {"Maggio", "Di"}, {"Maio", "Di"}, {"Mattina", "Di"}, {"Pietro", "Di"}, {"Sanzo", "Di"}, {"Roberto", "Dipiazza"}, {"Maurizio", "Dipietro"}, {"Daniela", "Dondi"}, {"Leonardo", "Donno"}, {"Giovanni", "Donzelli"}, {"Devis", "Dori"}, {"Marco", "Dreosto"}, {"Claudio", "Durigon"}, {"Meinhard", "Durnwalder"}, {"Michele", "Emiliano"}, {"Eleonora", "Evi"}, {"Alan", "Fabbri"}, {"Giuseppe", "Falcomatà"}, {"Anna", "Fallucchi"}, {"Davide", "Faraone"}, {"Marta", "Farolfi"}, {"Marta", "Fascina"}, {"Piero", "Fassino"}, {"Giovambattista", "Fazzolari"}, {"Claudio", "Fazzone"}, {"Giorgio", "Fede"}, {"Massimiliano", "Fedriga"}, {"Emiliano", "Fenu"}, {"Tullio", "Ferrante"}, {"Sara", "Ferrari"}, {"Wanda", "Ferro"}, {"Gianluca", "Festa"}, {"Francesco", "Filini"}, {"Michele", "Fina"}, {"Marco", "Fioravanti"}, {"Nicola", "Fiorita"}, {"Raffaele", "Fitto"}, {"Aurora", "Floridia"}, {"Barbara", "Floridia"}, {"Attilio", "Fontana"}, {"Ilaria", "Fontana"}, {"Lorenzo", "Fontana"}, {"Pietro", "Fontanini"}, {"Antonella", "Forattini"}, {"Paolo", "Formentini"}, {"Federico", "Fornaro"}, {"Emiliano", "Fossi"}, {"Tommaso", "Foti"}, {"Fabrizio", "Fracassi"}, {"Silvio", "Franceschelli"}, {"Dario", "Franceschini"}, {"Paola", "Frassinetti"}, {"Rebecca", "Frassini"}, {"Nicola", "Fratoianni"}, {"Silvia", "Fregolent"}, {"Federico", "Freni"}, {"Maria", "Frijia"}, {"Claudia", "Frontini"}, {"Maurizio", "Fugatti"}, {"Andrea", "Furegato"}, {"Marco", "Furfaro"}, {"Domenico", "Furgiuele"}, {"Annamaria", "Furlan"}, {"Maria", "Gadda"}, {"Edoardo", "Gaffeo"}, {"Gianluca", "Galimberti"}, {"Francesco", "Gallo"}, {"Roberto", "Gambino"}, {"Massimo", "Garavaglia"}, {"Elisabetta", "Gardini"}, {"Maurizio", "Gasparri"}, {"Giandiego", "Gatta"}, {"Mauro", "Gattinoni"}, {"Vannia", "Gava"}, {"Renate", "Gebhard"}, {"Matteo", "Gelmetti"}, {"Maria", "Gelmini"}, {"Marcello", "Gemmato"}, {"Antonino", "Germanà"}, {"Alessandro", "Ghinelli"}, {"Valentina", "Ghio"}, {"Francesca", "Ghirra"}, {"Andrea", "Giaccone"}, {"Roberto", "Giachetti"}, {"Francesco", "Giacobbe"}, {"Dario", "Giagoni"}, {"Federico", "Gianassi"}, {"Eugenio", "Giani"}, {"Vigna", "Giglio"}, {"Sergio", "Giordani"}, {"Antonio", "Giordano"}, {"Giancarlo", "Giorgetti"}, {"Carmen", "Giorgianni"}, {"Andrea", "Giorgis"}, {"Silvio", "Giovine"}, {"Gian", "Girelli"}, {"Carla", "Giuliano"}, {"Andrea", "Gnassi"}, {"Giorgio", "Gori"}, {"Roberto", "Gravina"}, {"Stefano", "Graziano"}, {"Claudia", "Gribaudo"}, {"Beppe", "Grillo"}, {"Marco", "Grimaldi"}, {"Valentina", "Grippo"}, {"Naike", "Gruppioni"}, {"Roberto", "Gualtieri"}, {"Michele", "Gubitosa"}, {"Lorenzo", "Guerini"}, {"Maria", "Guerra"}, {"Michele", "Guerra"}, {"Antonio", "Guidi"}, {"Barbara", "Guidolin"}, {"Alberto", "Gusmeroli"}, {"Giovanna", "Iacono"}, {"Dario", "Iaia"}, {"Franco", "Ianeselli"}, {"Antonio", "Iannone"}, {"Antonino", "Iaria"}, {"Massimiliano", "Iervolino"}, {"Igor", "Iezzi"}, {"Nicola", "Irto"}, {"Francesco", "Italia"}, {"Sara", "Kelany"}, {"Patty", "L'Abbate"}, {"Marca", "La"}, {"Pietra", "La"}, {"Porta", "La"}, {"Russa", "La"}, {"Salandra", "La"}, {"Marco", "Lacarra"}, {"Roberto", "Lagalla"}, {"Silvio", "Lai"}, {"Gianni", "Lampis"}, {"Elisabetta", "Lancellotta"}, {"Marcello", "Lanotte"}, {"Giorgia", "Latini"}, {"Leonardo", "Latini"}, {"Mauro", "Laus"}, {"Erik", "Lavevaz"}, {"Arianna", "Lazzarini"}, {"Maurizio", "Leo"}, {"Elena", "Leonardi"}, {"Matteo", "Lepore"}, {"Enrico", "Letta"}, {"Ettore", "Licheri"}, {"Sabrina", "Licheri"}, {"Mattia", "Limardo"}, {"Guido", "Liris"}, {"Marco", "Lisei"}, {"Russo", "Lo"}, {"Alessandra", "Locatelli"}, {"Simona", "Loizzo"}, {"Francesco", "Lollobrigida"}, {"Marco", "Lombardo"}, {"Arnaldo", "Lomuti"}, {"Eliana", "Longi"}, {"Emanuele", "Loperfido"}, {"Ada", "Lopreiato"}, {"Pietro", "Lorefice"}, {"Beatrice", "Lorenzin"}, {"Alberto", "Losacco"}, {"Claudio", "Lotito"}, {"Giorgio", "Lovecchio"}, {"Ylenja", "Lucaselli"}, {"Maurizio", "Lupi"}, {"Elena", "Maccanti"}, {"Carlo", "Maccari"}, {"Maria", "Madia"}, {"Novo", "Maerna"}, {"Gianpietro", "Maffoni"}, {"Riccardo", "Magi"}, {"Tino", "Magni"}, {"Giovanni", "Maiorano"}, {"Alessandra", "Maiorino"}, {"Lorenzo", "Malagola"}, {"Mauro", "Malaguti"}, {"Lucio", "Malan"}, {"Ilenia", "Malavasi"}, {"Simona", "Malpezzi"}, {"Patrizia", "Manassero"}, {"Daniele", "Manca"}, {"Valeria", "Mancinelli"}, {"Claudio", "Mancini"}, {"Paola", "Mancini"}, {"Franco", "Manes"}, {"Gaetano", "Manfredi"}, {"Giuseppe", "Mangialavori"}, {"Lucrezia", "Mantovani"}, {"Irene", "Manzi"}, {"Enzo", "Maraio"}, {"Luigi", "Marattin"}, {"Riccardo", "Marchetti"}, {"Aliprandi", "Marchetto"}, {"Silvia", "Marchionini"}, {"Paolo", "Mareschi"}, {"Francesco", "Mari"}, {"Maria", "Marino"}, {"Roberto", "Maroni"}, {"Patrizia", "Marrocco"}, {"Marco", "Marsilio"}, {"Andrea", "Martella"}, {"Roberto", "Marti"}, {"Bruno", "Marton"}, {"Andrea", "Mascaretti"}, {"Ciro", "Maschio"}, {"Carlo", "Masci"}, {"Clemente", "Mastella"}, {"Riccardo", "Mastrangeli"}, {"Domenico", "Matera"}, {"Mariangela", "Matera"}, {"Simonetta", "Matone"}, {"Nicole", "Matteoni"}, {"Aldo", "Mattia"}, {"Stefano", "Maullu"}, {"Matteo", "Mauri"}, {"Orfeo", "Mazzella"}, {"Erica", "Mazzetti"}, {"Gianmarco", "Mazzi"}, {"Filippo", "Melchiorre"}, {"Giorgia", "Meloni"}, {"Marco", "Meloni"}, {"Roberto", "Menia"}, {"Lavinia", "Mennuni"}, {"Rinaldo", "Menucci"}, {"Virginio", "Merola"}, {"Ignazio", "Messina"}, {"Manlio", "Messina"}, {"Francesco", "Miccichè"}, {"Gianfanco", "Miccichè"}, {"Francesco", "Michelotti"}, {"Giovanna", "Miele"}, {"Ester", "Mieli"}, {"Massimo", "Milani"}, {"Antonino", "Minardo"}, {"Tilde", "Minasi"}, {"Franco", "Mirabelli"}, {"Antonio", "Misiani"}, {"Riccardo", "Molinari"}, {"Federico", "Mollicone"}, {"Nicola", "Molteni"}, {"Augusta", "Montaruli"}, {"Elisa", "Montemagni"}, {"Mario", "Monti"}, {"Roberto", "Morassut"}, {"Alessandro", "Morelli"}, {"Daniela", "Morfino"}, {"Maddalena", "Morgante"}, {"Pietro", "Morittu"}, {"Jacopo", "Morrone"}, {"Giorgio", "Mulè"}, {"Francesco", "Mura"}, {"Elena", "Murelli"}, {"Dafne", "Musolino"}, {"Nello", "Musumeci"}, {"Gian", "Muzzarelli"}, {"Vincenzo", "Napoli"}, {"Giorgio", "Napolitano"}, {"Dario", "Nardella"}, {"Gaetano", "Nastri"}, {"Gisella", "Naturale"}, {"Luigi", "Nave"}, {"Raffaele", "Nevi"}, {"Antonio", "Nicita"}, {"Tiziana", "Nisini"}, {"Luciano", "Nobili"}, {"Vita", "Nocco"}, {"Carlo", "Nordio"}, {"Gianni", "Nuti"}, {"Mario", "Occhiuto"}, {"Roberto", "Occhiuto"}, {"Federica", "Onori"}, {"Matteo", "Orfini"}, {"Andrea", "Orlando"}, {"Anna", "Orrico"}, {"Andrea", "Orsini"}, {"Fausto", "Orsomarso"}, {"Marco", "Osnato"}, {"Andrea", "Ostellari"}, {"Nicola", "Ottaviani"}, {"Marco", "Padovani"}, {"Andrea", "Paganella"}, {"Nazario", "Pagano"}, {"Ubaldo", "Pagano"}, {"Raffaella", "Paita"}, {"Mattia", "Palazzi"}, {"Alessandro", "Palombi"}, {"Massimiliano", "Panizzut"}, {"Gianluigi", "Paragone"}, {"Sandro", "Parcaroli"}, {"Mario", "Pardini"}, {"Adriano", "Paroli"}, {"Dario", "Parrini"}, {"Giulia", "Pastorella"}, {"Luca", "Pastorino"}, {"Annarita", "Patriarca"}, {"Pietro", "Patton"}, {"Stefano", "Patuanelli"}, {"Emma", "Pavanelli"}, {"Roberto", "Pella"}, {"Marco", "Pellegrini"}, {"Andrea", "Pellicini"}, {"Vinicio", "Peluffo"}, {"Pasqualino", "Penza"}, {"Marcello", "Pera"}, {"Pierluigi", "Peracchini"}, {"Marco", "Perissa"}, {"Francesco", "Persiani"}, {"Giovanna", "Petrenga"}, {"Simona", "Petrucci"}, {"Renzo", "Piano"}, {"Matteo", "Piantedosi"}, {"Elisabetta", "Piccolotti"}, {"Fratin", "Pichetto"}, {"Attilio", "Pierro"}, {"Fabio", "Pietrella"}, {"Simone", "Pillon"}, {"Paolo", "Pilotto"}, {"Luca", "Pirondini"}, {"Daisy", "Pirovano"}, {"Elisa", "Pirro"}, {"Calogero", "Pisano"}, {"Pietro", "Pittalis"}, {"Graziano", "Pizzimenti"}, {"Salvo", "Pogliese"}, {"Caria", "Polidori"}, {"Barbara", "Polo"}, {"Fabio", "Porta"}, {"Manfredi", "Potenti"}, {"Emanuele", "Pozzolo"}, {"Erik", "Pretto"}, {"Emanuele", "Prisco"}, {"Romano", "Prodi"}, {"Giuseppe", "Provenzano"}, {"Stefania", "Pucciarelli"}, {"Paolo", "Pulciani"}, {"Procopio", "Quartapelle"}, {"Andrea", "Quartini"}, {"Angela", "Raffa"}, {"Virginia", "Raggi"}, {"Carmine", "Raimondo"}, {"Fabio", "Rampelli"}, {"Vincenza", "Rando"}, {"Ernesto", "Rapani"}, {"Alessandro", "Rapinese"}, {"Maurizio", "Rasero"}, {"Sergio", "Rastrelli"}, {"Isabella", "Rauti"}, {"Laura", "Ravetto"}, {"Matteo", "Renzi"}, {"Matteo", "Ricci"}, {"Marianna", "Ricciardi"}, {"Riccardo", "Ricciardi"}, {"Toni", "Ricciardi"}, {"Matteo", "Richetti"}, {"Edoardo", "Rixi"}, {"Walter", "Rizzetto"}, {"Marco", "Rizzo"}, {"Eugenia", "Roccella"}, {"Silvia", "Roggiani"}, {"Tatjana", "Rojc"}, {"Francesco", "Romano"}, {"Massimiliano", "Romeo"}, {"Andrea", "Romizi"}, {"Licia", "Ronzulli"}, {"Gianni", "Rosa"}, {"Ettore", "Rosato"}, {"Fabio", "Roscani"}, {"Cristina", "Rossello"}, {"Andrea", "Rossi"}, {"Angelo", "Rossi"}, {"Fabrizio", "Rossi"}, {"Riccardo", "Rossi"}, {"Matteo", "Rosso"}, {"Roberto", "Rosso"}, {"Anna", "Rossomando"}, {"Mauro", "Rotelli"}, {"Gianfranco", "Rotondi"}, {"Francesco", "Rubano"}, {"Carlo", "Rubbia"}, {"Francesco", "Rucco"}, {"Daniela", "Ruffino"}, {"Massimo", "Ruspandini"}, {"Gaetana", "Russo"}, {"Marco", "Russo"}, {"Paolo", "Russo"}, {"Raoul", "Russo"}, {"Jotti", "Saccani"}, {"Jamil", "Sadegholvaad"}, {"Fabrizio", "Sala"}, {"Giuseppe", "Sala"}, {"Salvatore", "Sallemi"}, {"Carlo", "Salvemini"}, {"Luca", "Salvetti"}, {"Matteo", "Salvini"}, {"Giorgio", "Salvitti"}, {"Gennaro", "Sangiuliano"}, {"Massimiliano", "Sanna"}, {"Daniela", "Santanchè"}, {"Agostino", "Santillo"}, {"Marco", "Sarracino"}, {"Rossano", "Sasso"}, {"Giovanni", "Satta"}, {"Luca", "Sbardella"}, {"Daniela", "Sbrollini"}, {"Claudio", "Scajola"}, {"Ivan", "Scalfarotto"}, {"Marco", "Scaramellini"}, {"Rachele", "Scarpa"}, {"Roberto", "Scarpinato"}, {"Filippo", "Scerra"}, {"Di", "Schiano"}, {"Renato", "Schifani"}, {"Marta", "Schifone"}, {"Orazio", "Schillaci"}, {"Elly", "Schlein"}, {"Manfred", "Schullian"}, {"Arturo", "Scotto"}, {"Marco", "Scurria"}, {"Elisa", "Scutellà"}, {"Liliana", "Segre"}, {"Martina", "Semenzato"}, {"Debora", "Serracchiani"}, {"Vittorio", "Sgarbi"}, {"Etelwardo", "Sigismondi"}, {"Francesco", "Silvestri"}, {"Rachele", "Silvestri"}, {"Francesco", "Silvestro"}, {"Marco", "Silvestroni"}, {"Marco", "Simiani"}, {"Daniele", "Sinibaldi"}, {"Matilde", "Siracusano"}, {"Elena", "Sironi"}, {"Sandro", "Sisler"}, {"Francesco", "Sisto"}, {"Andrea", "Soddu"}, {"Christian", "Solinas"}, {"Alessandro", "Sorte"}, {"Giulio", "Sottanelli"}, {"Aboubakar", "Soumahoro"}, {"Luigi", "Spagnolli"}, {"Nicoletta", "Spelgatti"}, {"Roberto", "Speranza"}, {"Raffaele", "Speranzon"}, {"Domenica", "Spinelli"}, {"Gilda", "Sportiello"}, {"Luca", "Squeri"}, {"Claudio", "Stefanazzi"}, {"Alberto", "Stefani"}, {"Erika", "Stefani"}, {"Dieter", "Steger"}, {"Nicola", "Stumpo"}, {"Valeria", "Sudano"}, {"Bruno", "Tabacci"}, {"Antonio", "Tajani"}, {"Katia", "Tarasconi"}, {"Rosaria", "Tassinari"}, {"Chiara", "Tenerini"}, {"di", "Terzi"}, {"Donatella", "Tesei"}, {"Guerino", "Testa"}, {"Elena", "Testor"}, {"Franco", "Tirelli"}, {"Luca", "Toccalini"}, {"Alessandra", "Todde"}, {"Donato", "Toma"}, {"Alessandro", "Tomasi"}, {"Damiano", "Tommasi"}, {"Danilo", "Toninelli"}, {"Daniela", "Torto"}, {"Paolo", "Tosato"}, {"Flavio", "Tosi"}, {"Giovanni", "Toti"}, {"Paolo", "Trancassini"}, {"Giacomo", "Tranchida"}, {"Roberto", "Traversi"}, {"Andrea", "Tremaglia"}, {"Giulio", "Tremonti"}, {"Antonio", "Trevisi"}, {"Paolo", "Truzzu"}, {"Francesca", "Tubetti"}, {"Riccardo", "Tucci"}, {"Mario", "Turco"}, {"Julia", "Unterberger"}, {"Adolfo", "Urso"}, {"Alessandro", "Urzì"}, {"Stefano", "Vaccari"}, {"Giuseppe", "Valditara"}, {"Valeria", "Valenta"}, {"Carmine", "Valente"}, {"Maria", "Varchi"}, {"Antonello", "Velardi"}, {"Walter", "Veltroni"}, {"Francesco", "Verducci"}, {"Walter", "Verini"}, {"Giusy", "Versace"}, {"Imma", "Vietri"}, {"Gianluca", "Vinci"}, {"Colonna", "Vivarelli"}, {"Vincenzo", "Voce"}, {"Andrea", "Volpi"}, {"Francesco", "Zaffini"}, {"Luca", "Zaia"}, {"Ylenia", "Zambito"}, {"Sandra", "Zampa"}, {"Alessandro", "Zan"}, {"Luana", "Zanella"}, {"Pierantonio", "Zanettin"}, {"Paolo", "Zangrillo"}, {"Giorgio", "Zanni"}, {"Filiberto", "Zaratti"}, {"Gian", "Zattini"}, {"Antonella", "Zedda"}, {"Rodolfo", "Ziberna"}, {"Edoardo", "Ziello"}, {"Nicola", "Zingaretti"}, {"Gianpiero", "Zinzi"}, {"Eugenio", "Zoffili"}, {"Riccardo", "Zucconi"}, {"Ignazio", "Zullo"}, {"Immacolata", "Zurzolo"}, {"(", ""}}
