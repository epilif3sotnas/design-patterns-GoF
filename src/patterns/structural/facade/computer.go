package facade


type Computer struct {
	cpu *Cpu
	hardDrive *HardDrive
	memory *Memory
}


func NewComputer() *Computer {
	return &Computer{
		cpu: NewCpu(),
		hardDrive: NewHardDrive(),
		memory: NewMemory(),
	}
}

func (self *Computer) Start() {
	self.cpu.Freeze()
	self.memory.Load(0x80000004, self.hardDrive.Read(0x8aaf04, 16_000_000))
	self.cpu.Jump(0x80fef1)
	self.cpu.Execute()
}

func (self *Computer) GetCpu() *Cpu {
	return self.cpu
}

func (self *Computer) SetCpu(cpu *Cpu) {
	self.cpu = cpu
}

func (self *Computer) GetHardDrive() *HardDrive {
	return self.hardDrive
}

func (self *Computer) SetHardDrive(hardDrive *HardDrive) {
	self.hardDrive = hardDrive
}

func (self *Computer) GetMemory() *Memory {
	return self.memory
}

func (self *Computer) SetMemory(memory *Memory) {
	self.memory = memory
}