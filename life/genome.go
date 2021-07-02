package life

type Genome struct {
	physiology PhysiologyGenome
	behaviour  BehaviourGenome
}

type PhysiologyGenome struct {
	dna []byte
}

type BehaviourGenome struct {
	dna []byte
}
