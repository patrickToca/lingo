// generated by stringer -type=feature -output=features_string.go; DO NOT EDIT

package dep

import "fmt"

const _feature_name = "s0ws1ws2wb0wb1wb2ws0l1ws0r1ws0l2ws0r2ws0llws0rrws1l1ws1r1ws1l2ws1r2ws1llws1rrws0ts1ts2tb0tb1tb2ts0l1ts0r1ts0l2ts0r2ts0llts0rrts1l1ts1r1ts1l2ts1r2ts1llts1rrts0l1ds0r1ds0l2ds0r2ds0llds0rrds1l1ds1r1ds1l2ds1r2ds1llds1rrdMAXFEATURE"

var _feature_index = [...]uint8{0, 3, 6, 9, 12, 15, 18, 23, 28, 33, 38, 43, 48, 53, 58, 63, 68, 73, 78, 81, 84, 87, 90, 93, 96, 101, 106, 111, 116, 121, 126, 131, 136, 141, 146, 151, 156, 161, 166, 171, 176, 181, 186, 191, 196, 201, 206, 211, 216, 226}

func (i feature) String() string {
	if i < 0 || i >= feature(len(_feature_index)-1) {
		return fmt.Sprintf("feature(%d)", i)
	}
	return _feature_name[_feature_index[i]:_feature_index[i+1]]
}