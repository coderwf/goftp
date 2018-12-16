package manager

//目录管理

type Dir struct {
	cwd           string   //当前目录
}

func (d *Dir) Pwd() string{
	return d.cwd
}

func (d *Dir) Cd() string{

}

func (d *Dir) Mkdir(){

}

