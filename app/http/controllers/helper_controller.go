package controllers

type HelperController struct {
}

func NewHelperController() *HelperController {
	return &HelperController{
		//Inject services
	}
}

func (r *HelperController) FindStringIntoSlice(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
