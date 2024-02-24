package rs485

import . "github.com/volkszaehler/mbmd/meters"

func init() {
	Register("SDM630", NewSDM630Producer)
}

type SDM630Producer struct {
	Opcodes
}

func NewSDM630Producer() Producer {
	/**
	 * Opcodes as defined by Eastron SDM630.
	 * See http://bg-etech.de/download/manual/SDM630Register.pdf
	 * This is to a large extent a superset of all SDM devices, however there are
	 * subtle differences (see 220, 230). Some opcodes might not work on some devices.
	 */
	ops := Opcodes{
		// VoltageL1:     		0x0000, // Voltage L1 [V]
		// VoltageL2:     		0x0002, // Voltage L2 [V]
		// VoltageL3:     		0x0004, // Voltage L3 [V]
		CurrentL1:				0x0006, // Current L1 [A]
		CurrentL2:				0x0008, // Current L2 [A]
		CurrentL3:				0x000A, // Current L3 [A]
		PowerL1:  				0x000C, // Active Power L1 [W]
		PowerL2:  				0x000E, // Active Power L2 [W]
		PowerL3:  				0x0010, // Active Power L3 [W]
		Power:    				0x0034, // Total Active Power [W]
		// ApparentPower: 		0x0038, // Total Apparent Power [VA]
		ReactivePowerL1: 		0x0018, // Reactive Power L1 [VAr]
		ReactivePowerL2: 		0x001a, // Reactive Power L2 [VAr]
		ReactivePowerL3: 		0x001c, // Reactive Power L3 [VAr]
		ReactivePower:   		0x003C, // Total Reactive Power [VAr]
		// ImportPower: 		0x0054, // Import Power [W]
		ImportL1:      			0x015a, // Import Energy L1 [kWh]
		ImportL2:      			0x015c, // Import Energy L2 [kWh]
		ImportL3:      			0x015e, // Import Energy L3 [kWh]
		Import: 				0x0048, // Import Energy [kWh]
		// ExportL1:      		0x0160, // Export Energy L1 [kWh]
		// ExportL2:      		0x0162, // Export Energy L2 [kWh]
		// ExportL3:      		0x0164, // Export Energy L3 [kWh]
		// Export: 				0x004a, // Export Energy [kWh]
		// SumL1:				0x0166, // Sum -> import + export Energy L1 [kWh]
		// SumL2:         		0x0168, // Sum -> import + export Energy L2 [kWh]
		// SumL3:         		0x016a, // Sum -> import + export Energy L3 [kWh]
		// Sum:				    0x0156, // Sum Energy L1+L2+L3 [kWh]
		// CosphiL1:      		0x001e, // Power Factor L1
		// CosphiL2:      		0x0020, // Power Factor L2
		// CosphiL3:      		0x0022, // Power Factor L3
		// Cosphi:        		0x003e, // Total system Power Factor
		// THDL1:         		0x00ea, // Phase 1 L/N volts THD [%]
		// THDL2:         		0x00ec, // Phase 2 L/N volts THD [%]
		// THDL3:         		0x00ee, // Phase 3 L/N volts THD [%]
		// THD:           		0x00F8, // Average line to neutral volts THD [%]
		// Frequency:     		0x0046, // Line Frequency [Hz]
		// L1THDCurrent: 		0x00F0, // Phase 1 Current THD [%]
		// L2THDCurrent: 		0x00F2, // Phase 2 Current THD [%]
		// L3THDCurrent: 		0x00F4, // Phase 3 Current THD [%]
		// AvgTHDCurrent: 		0x00Fa, // Average line current THD [%]
		// ApparentImportPower: 0x0064, // Total system VA demand [VA]
	}
	return &SDM630Producer{Opcodes: ops}
}

func (p *SDM630Producer) Description() string {
	return "Eastron SDM630"
}

func (p *SDM630Producer) snip(iec Measurement) Operation {
	operation := Operation{
		FuncCode:  ReadInputReg,
		OpCode:    p.Opcode(iec),
		ReadLen:   2,
		IEC61850:  iec,
		Transform: RTUIeee754ToFloat64,
	}
	return operation
}

func (p *SDM630Producer) Probe() Operation {
	return p.snip(Voltage)
}

func (p *SDM630Producer) Produce() (res []Operation) {
	for op := range p.Opcodes {
		res = append(res, p.snip(op))
	}

	return res
}
