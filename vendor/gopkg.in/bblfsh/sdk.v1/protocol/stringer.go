// Code generated by "stringer -type=Status,Encoding -output stringer.go"; DO NOT EDIT.

package protocol

import "strconv"

const _Status_name = "OkErrorFatal"

var _Status_index = [...]uint8{0, 2, 7, 12}

func (i Status) String() string {
	if i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}

const _Encoding_name = "UTF8Base64"

var _Encoding_index = [...]uint8{0, 4, 10}

func (i Encoding) String() string {
	if i >= Encoding(len(_Encoding_index)-1) {
		return "Encoding(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Encoding_name[_Encoding_index[i]:_Encoding_index[i+1]]
}