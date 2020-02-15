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

func (pkgs *Packages) Remove(pkg *Package) *Packages {
	filtered := Packages{}
	for _, p := range *pkgs {
		if p.equals(pkg) {
			continue
		}
		filtered = append(filtered, p)
	}
	*pkgs = filtered
	return pkgs
}

func (p *Package) equals(other *Package) bool {
	return p.Manager == other.Manager &&
		p.Name == other.Name
}
