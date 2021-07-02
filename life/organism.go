package life

import "reality.rehab/catastrophe/world"

const (
	MAX_ORG_HEIGHT = 25
	MAX_ORG_WIDTH  = 25
)

type Organism struct {
	genome Genome
	form   world.Object
}

func New(g Genome, w world.World) *Organism {
	o := &Organism{genome: g}
	o.form = spawn(g.physiology)
	return o
}

// So, here we translate the PhysiologyGenome DNA into something that can be plotted
func spawn(p PhysiologyGenome) *Object {
	obj := world.CreateObject(MAX_ORG_WIDTH, MAX_ORG_HEIGHT)

	for i, c := range p.dna {
	}

	obj.Trim()

	return f
}
