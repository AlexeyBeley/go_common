packkge logger

import fmt
type Logger struct{

}

func (l Logger) r(str string, args ...any){
fmt.Print("%v, %v",str, args)
}
