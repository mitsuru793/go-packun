package pkg

type Package struct {
	Manager string
	Name    string
	Tags    []string
}

type Packages []Package

func (pkgs *Packages) Exists(newPkg *Package) bool {
	for _, p := range *pkgs {
		if p.equals(newPkg) {
			return true
		}
	}
	return false
}

func (pkgs *Packages) Add(newPkg *Package) *Packages {
	for _, p := range *pkgs {
		if p.equals(newPkg) {
			return pkgs
		}
	}
	*pkgs = append(*pkgs, *newPkg)
	return pkgs
}

func (p *Package) equals(other *Package) bool {
	return p.Manager == other.Manager &&
		p.Name == other.Name
}
