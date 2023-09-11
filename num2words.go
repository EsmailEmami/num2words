package num2words

import "math"

// how many digit's groups to process
const groupsNumber int = 6

var _smallNumbers = []string{
	"صفر", "یک", "دو", "سه", "چهار",
	"پنج", "شش", "هفت", "هشت", "نه",
	"ده", "یازده", "دوازده", "سیزده", "چهارده",
	"پانزده", "شانزده", "هفده", "هجده", "نوزده",
}
var _tens = []string{
	"", "", "بیست", "سی", "چهل", "پنجاه",
	"شصت", "هفتاد", "هشتاد", "نود",
}

var _hundreds = []string{
	"", "یکصد", "دویست", "سیصد", "چهارصد", "پانصد",
	"ششصد", "هفتصد", "هشتصد", "نهصد",
}

var _scaleNumbers = []string{
	"", "هزار", "میلیون", "میلیارد", "بیلیون", "بیلیارد",
}

type digitGroup int

// Convert converts number into the words representation.
func Convert(number int) string {
	// Zero rule
	if number == 0 {
		return _smallNumbers[0]
	}

	// Divide into three-digits group
	var groups [groupsNumber]digitGroup
	positive := math.Abs(float64(number))

	// Form three-digit groups
	for i := 0; i < groupsNumber; i++ {
		groups[i] = digitGroup(math.Mod(positive, 1000))
		positive /= 1000
	}

	var textGroup [groupsNumber]string
	for i := 0; i < groupsNumber; i++ {
		textGroup[i] = digitGroup2Text(groups[i])
	}
	combined := textGroup[0]

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			prefix := textGroup[i] + " " + _scaleNumbers[i]

			if len(combined) != 0 {
				prefix += separator()
			}

			combined = prefix + combined
		}
	}

	if number < 0 {
		combined = "منهای " + combined
	}

	return combined
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func digitGroup2Text(group digitGroup) (ret string) {
	hundreds := group / 100
	tensUnits := intMod(int(group), 100)

	if hundreds != 0 {
		ret += _hundreds[hundreds]

		if tensUnits != 0 {
			ret += separator()
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		ret += _tens[tens]

		if units != 0 {
			ret += separator() + _smallNumbers[units]
		}
	} else if tensUnits != 0 {
		ret += _smallNumbers[tensUnits]
	}

	return
}

// separator returns proper separator string between
// number groups.
func separator() string {
	return " و "
}

func init() {

}
