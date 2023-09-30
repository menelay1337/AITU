package main
import "fmt"



type doctor interface {
	addPatient(patient)
	deletePatient()
	notifyPatients()
}

type therapist struct{
	name string
	surname string
	patients []patient
}

func (d *therapist) addPatient(p patient){
	d.patients = append(d.patients, p)
}

func (d *therapist) deletePatient(p patient){
	for idx, p1 := range d.patients {
		if p1 == p {
			d.patients = append(d.patients[0:idx], d.patients[idx+1:]...)
		}
	}
}

func (d *therapist) notifyPatients(){
	for _, pat := range d.patients {
		pat.notify(d.name + " " + d.surname, "Therapist")
	}
}


type patient struct {
	name string
	surname string
	ID uint 
}

func (p *patient) notify(doctorFullName string, doctorSpecification string) {	
	fmt.Printf("'doctor %s %s is now has entry!' email sended to patient %s with id: %d\n", doctorFullName, doctorSpecification, p.name + " " + p.surname, p.ID)
}

func main(){
	pat1 := patient{"Maksat", "Zhalkysbaev", 132465}
	pat2 := patient{"Zhalgas", "Konebayev", 123423}
	pat3 := patient{"Maxim", "Zaycev", 345123}

	doc1 := therapist{"Kanat", "Kairanas", make([]patient, 0, 10)}
	
	doc1.addPatient(pat1)
	doc1.addPatient(pat2)

	fmt.Println(pat3)
	doc1.notifyPatients()
	fmt.Println()


	doc1.deletePatient(pat2)
	
	doc1.addPatient(pat3)
	fmt.Println(doc1.patients)
	doc1.notifyPatients()


}
