package utils

// Int32Set
type Int32Set struct {
	Set map[int32]struct{}
}

func NewInt32Set() *Int32Set {
	s := Int32Set{}
	s.Set = make(map[int32]struct{})
	return &s
}

func (s *Int32Set) InitFromSlice(inp []int32) *Int32Set {
	s.Set = make(map[int32]struct{}, len(inp))
	for _, item := range inp {
		s.Add(item)
	}
	return s
}

func (s *Int32Set) Copy() *Int32Set {
	return NewInt32Set().InitFromSlice(s.GetSlice())
}

func (s *Int32Set) check() {
	if s.Set == nil {
		s.Set = make(map[int32]struct{})
	}
}

func (s *Int32Set) Length() int {
	s.check()
	return len(s.Set)
}

func (s *Int32Set) Add(inp int32) bool {
	s.check()
	_, found := s.Set[inp]
	if !found {
		s.Set[inp] = struct{}{}
	}
	return !found //False if exist already
}

func (s *Int32Set) Insert(inp int32) *Int32Set {
	s.Add(inp)
	return s
}

func (s *Int32Set) Update(inp []int32) {
	s.check()
	for _, item := range inp {
		s.Add(item)
	}
}

func (s *Int32Set) Delete(inp int32) bool {
	s.check()
	_, found := s.Set[inp]
	if found {
		delete(s.Set, inp)
	}
	return found //False if not exist already
}

func (s *Int32Set) Remove(inp int32) *Int32Set {
	s.Delete(inp)
	return s
}

func (s *Int32Set) GetSlice() []int32 {
	s.check()
	outp := []int32{}
	for k := range s.Set {
		outp = append(outp, k)
	}
	return outp
}

func (s *Int32Set) Contains(inp int32) bool {
	_, ok := s.Set[inp]
	return ok
}

func (s *Int32Set) Equals(inp *Int32Set) bool {
	return (s.Length() == 0 && inp.Length() == 0) ||
		(s.Length() == inp.Length() &&
			s.Copy().Not(inp).Length() == 0)
}

func (s *Int32Set) Merge(inp *Int32Set) {
	s.check()
	for k := range inp.Set {
		s.Set[k] = struct{}{}
	}
}

func (s *Int32Set) Combine(inp *Int32Set) *Int32Set {
	s.check()
	combination := s.Copy()
	for k := range inp.Set {
		combination.Set[k] = struct{}{}
	}
	return combination
}

func (s *Int32Set) And(s2 *Int32Set) *Int32Set {
	s.check()
	intersection := NewInt32Set()
	var compareSet *Int32Set
	var iterateSet *Int32Set
	if s2 == nil {
		return intersection
	}
	// Iterate the shorter set.
	if s2.Length() > s.Length() {
		compareSet = s2
		iterateSet = s
	} else {
		compareSet = s
		iterateSet = s2
	}
	for _, item := range iterateSet.GetSlice() {
		if compareSet.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s *Int32Set) Not(s2 *Int32Set) *Int32Set {
	s.check()
	if s2 == nil {
		return s.Copy()
	}
	notIds := NewInt32Set()
	for item := range s.Set {
		if !s2.Contains(item) {
			notIds.Add(item)
		}
	}
	return notIds
}

// Int64Set
type Int64Set struct {
	Set map[int64]struct{}
}

func NewInt64Set() *Int64Set {
	s := Int64Set{}
	s.Set = make(map[int64]struct{})
	return &s
}

func (s *Int64Set) InitFromSlice(inp []int64) *Int64Set {
	s.Set = make(map[int64]struct{}, len(inp))
	for _, item := range inp {
		s.Add(item)
	}
	return s
}

func (s *Int64Set) Copy() *Int64Set {
	return NewInt64Set().InitFromSlice(s.GetSlice())
}

func (s *Int64Set) check() {
	if s.Set == nil {
		s.Set = make(map[int64]struct{})
	}
}

func (s *Int64Set) Length() int {
	s.check()
	return len(s.Set)
}

func (s *Int64Set) Add(inp int64) bool {
	s.check()
	_, found := s.Set[inp]
	if !found {
		s.Set[inp] = struct{}{}
	}
	return !found //False if exist already
}

func (s *Int64Set) Insert(inp int64) *Int64Set {
	s.Add(inp)
	return s
}

func (s *Int64Set) Update(inp []int64) {
	s.check()
	for _, item := range inp {
		s.Add(item)
	}
}

func (s *Int64Set) Delete(inp int64) bool {
	s.check()
	_, found := s.Set[inp]
	if found {
		delete(s.Set, inp)
	}
	return found //False if not exist already
}

func (s *Int64Set) Remove(inp int64) *Int64Set {
	s.Delete(inp)
	return s
}

func (s *Int64Set) GetSlice() []int64 {
	s.check()
	outp := []int64{}
	for k := range s.Set {
		outp = append(outp, k)
	}
	return outp
}

func (s *Int64Set) Contains(inp int64) bool {
	_, ok := s.Set[inp]
	return ok
}

func (s *Int64Set) Equals(inp *Int64Set) bool {
	return (s.Length() == 0 && inp.Length() == 0) ||
		(s.Length() == inp.Length() &&
			s.Copy().Not(inp).Length() == 0)
}

func (s *Int64Set) Merge(inp *Int64Set) {
	s.check()
	for k := range inp.Set {
		s.Set[k] = struct{}{}
	}
}

func (s *Int64Set) Combine(inp *Int64Set) *Int64Set {
	s.check()
	combination := s.Copy()
	for k := range inp.Set {
		combination.Set[k] = struct{}{}
	}
	return combination
}

func (s *Int64Set) And(s2 *Int64Set) *Int64Set {
	s.check()
	intersection := NewInt64Set()
	var compareSet *Int64Set
	var iterateSet *Int64Set
	if s2 == nil {
		return intersection
	}
	// Iterate the shorter set.
	if s2.Length() > s.Length() {
		compareSet = s2
		iterateSet = s
	} else {
		compareSet = s
		iterateSet = s2
	}
	for _, item := range iterateSet.GetSlice() {
		if compareSet.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s *Int64Set) Not(s2 *Int64Set) *Int64Set {
	s.check()
	if s2 == nil {
		return s.Copy()
	}
	notIds := NewInt64Set()
	for item := range s.Set {
		if !s2.Contains(item) {
			notIds.Add(item)
		}
	}
	return notIds
}

// StringSet
type StringSet struct {
	Set map[string]struct{}
}

func NewStringSet() *StringSet {
	s := StringSet{}
	s.Set = make(map[string]struct{})
	return &s
}

func (s *StringSet) InitFromSlice(inp []string) *StringSet {
	s.Set = make(map[string]struct{}, len(inp))
	for _, item := range inp {
		s.Add(item)
	}
	return s
}

func (s *StringSet) Copy() *StringSet {
	return NewStringSet().InitFromSlice(s.GetSlice())
}

func (s *StringSet) check() {
	if s.Set == nil {
		s.Set = make(map[string]struct{})
	}
}

func (s *StringSet) Length() int {
	s.check()
	return len(s.Set)
}

func (s *StringSet) Add(inp string) bool {
	s.check()
	_, found := s.Set[inp]
	if !found {
		s.Set[inp] = struct{}{}
	}
	return !found //False if not exist already
}

func (s *StringSet) Insert(inp string) *StringSet {
	s.Add(inp)
	return s
}

func (s *StringSet) Update(inp []string) {
	s.check()
	for _, item := range inp {
		s.Add(item)
	}
}

func (s *StringSet) Delete(inp string) bool {
	s.check()
	_, found := s.Set[inp]
	if found {
		delete(s.Set, inp)
	}
	return found //False if exist already
}

func (s *StringSet) Remove(inp string) *StringSet {
	s.Delete(inp)
	return s
}

func (s *StringSet) GetSlice() []string {
	s.check()
	outp := []string{}
	for k := range s.Set {
		outp = append(outp, k)
	}
	return outp
}

func (s *StringSet) Contains(inp string) bool {
	_, ok := s.Set[inp]
	return ok
}

func (s *StringSet) Equals(inp *StringSet) bool {
	return (s.Length() == 0 && inp.Length() == 0) ||
		(s.Length() == inp.Length() &&
			s.Copy().Not(inp).Length() == 0)
}

func (s *StringSet) Merge(inp *StringSet) {
	s.check()
	for k := range inp.Set {
		s.Set[k] = struct{}{}
	}
}

func (s *StringSet) Combine(inp *StringSet) *StringSet {
	s.check()
	combination := s.Copy()
	for k := range inp.Set {
		combination.Set[k] = struct{}{}
	}
	return combination
}

func (s *StringSet) And(s2 *StringSet) *StringSet {
	s.check()
	intersection := NewStringSet()
	var compareSet *StringSet
	var iterateSet *StringSet
	if s2 == nil {
		return intersection
	}
	// Iterate the shorter set.
	if s2.Length() > s.Length() {
		compareSet = s2
		iterateSet = s
	} else {
		compareSet = s
		iterateSet = s2
	}
	for _, item := range iterateSet.GetSlice() {
		if compareSet.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s *StringSet) Not(s2 *StringSet) *StringSet {
	s.check()
	if s2 == nil {
		return s.Copy()
	}
	notIds := NewStringSet()
	for item := range s.Set {
		if !s2.Contains(item) {
			notIds.Add(item)
		}
	}
	return notIds
}

// MapStringSet
type MapStringSet struct {
	Sets map[string]*StringSet
}

func NewMapStringSet() *MapStringSet {
	return &MapStringSet{}
}

func (ms *MapStringSet) check() {
	if ms.Sets == nil {
		ms.Sets = make(map[string]*StringSet)
	}
}

func (ms *MapStringSet) checkKey(key string) {
	ms.check()
	if ms.Sets[key] == nil {
		ms.Sets[key] = NewStringSet()
	}
}

func (ms *MapStringSet) Length() int {
	ms.check()
	return len(ms.Sets)
}

func (ms *MapStringSet) Add(key, inp string) {
	ms.checkKey(key)
	ms.Sets[key].Add(inp)
}

func (ms *MapStringSet) GetSlice(key string) []string {
	ms.checkKey(key)
	if ms.Sets[key] == nil {
		return []string{}
	}
	return ms.Sets[key].GetSlice()
}

func (ms *MapStringSet) GetMap() map[string]*StringSet {
	ms.check()
	return ms.Sets
}

func (ms *MapStringSet) GetMapWithSlices() map[string][]string {
	withSlices := make(map[string][]string)
	for key := range ms.Sets {
		withSlices[key] = ms.Sets[key].GetSlice()
	}
	return withSlices
}

func (ms *MapStringSet) Contains(key, inp string) bool {
	ms.checkKey(key)
	return ms.GetMap()[key].Contains(inp)
}

func (ms *MapStringSet) Unwind() *StringSet {
	ms.check()
	items := NewStringSet()
	for _, item := range ms.Sets {
		items.Update(item.GetSlice())
	}
	return items
}

func (ms *MapStringSet) And(ms2 *MapStringSet) *StringSet {
	ms.check()
	intersection := NewStringSet()
	var compareSet *StringSet
	var iterateSet *StringSet
	if ms2 == nil {
		return intersection
	}
	// Iterate the shorter set.
	if ms2.Length() > ms.Length() {
		compareSet = ms2.Unwind()
		iterateSet = ms.Unwind()
	} else {
		compareSet = ms.Unwind()
		iterateSet = ms2.Unwind()
	}
	for _, item := range iterateSet.GetSlice() {
		if compareSet.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s *MapStringSet) Not(s2 *MapStringSet) *StringSet {
	s.check()
	compareA := s.Unwind()
	if s2 == nil {
		return compareA
	}
	notIds := NewStringSet()
	compareB := s2.Unwind()
	for item := range compareA.Set {
		if !compareB.Contains(item) {
			notIds.Add(item)
		}
	}
	return notIds
}

// InterfaceSet
type InterfaceSet struct {
	Set map[interface{}]struct{}
}

func NewInterfaceSet() *InterfaceSet {
	s := InterfaceSet{}
	s.Set = make(map[interface{}]struct{})
	return &s
}

func (s *InterfaceSet) InitFromSlice(inp []interface{}) *InterfaceSet {
	s.Set = make(map[interface{}]struct{}, len(inp))
	for _, item := range inp {
		s.Add(item)
	}
	return s
}

func (s *InterfaceSet) Copy() *InterfaceSet {
	return NewInterfaceSet().InitFromSlice(s.GetSlice())
}

func (s *InterfaceSet) check() {
	if s.Set == nil {
		s.Set = make(map[interface{}]struct{})
	}
}

func (s *InterfaceSet) Length() int {
	s.check()
	return len(s.Set)
}

func (s *InterfaceSet) Add(inp interface{}) bool {
	s.check()
	_, found := s.Set[inp]
	if !found {
		s.Set[inp] = struct{}{}
	}
	return !found //False if exist already
}

func (s *InterfaceSet) Insert(inp interface{}) *InterfaceSet {
	s.Add(inp)
	return s
}

func (s *InterfaceSet) Update(inp []interface{}) {
	s.check()
	for _, item := range inp {
		s.Add(item)
	}
}

func (s *InterfaceSet) Delete(inp interface{}) bool {
	s.check()
	_, found := s.Set[inp]
	if found {
		delete(s.Set, inp)
	}
	return found //False if not exist already
}

func (s *InterfaceSet) Remove(inp interface{}) *InterfaceSet {
	s.Delete(inp)
	return s
}

func (s *InterfaceSet) GetSlice() []interface{} {
	s.check()
	outp := make([]interface{}, 0)
	for k := range s.Set {
		outp = append(outp, k)
	}
	return outp
}

func (s *InterfaceSet) Contains(inp interface{}) bool {
	_, ok := s.Set[inp]
	return ok
}

func (s *InterfaceSet) Equals(inp *InterfaceSet) bool {
	return (s.Length() == 0 && inp.Length() == 0) ||
		(s.Length() == inp.Length() &&
			s.Copy().Not(inp).Length() == 0)
}

func (s *InterfaceSet) Merge(inp *InterfaceSet) {
	s.check()
	for k := range inp.Set {
		s.Set[k] = struct{}{}
	}
}

func (s *InterfaceSet) Combine(inp *InterfaceSet) *InterfaceSet {
	s.check()
	combination := s.Copy()
	for k := range inp.Set {
		combination.Set[k] = struct{}{}
	}
	return combination
}

func (s *InterfaceSet) And(s2 *InterfaceSet) *InterfaceSet {
	s.check()
	intersection := NewInterfaceSet()
	var compareSet *InterfaceSet
	var iterateSet *InterfaceSet
	if s2 == nil {
		return intersection
	}
	// Iterate the shorter set.
	if s2.Length() > s.Length() {
		compareSet = s2
		iterateSet = s
	} else {
		compareSet = s
		iterateSet = s2
	}
	for _, item := range iterateSet.GetSlice() {
		if compareSet.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s *InterfaceSet) Not(s2 *InterfaceSet) *InterfaceSet {
	s.check()
	if s2 == nil {
		return s.Copy()
	}
	notIds := NewInterfaceSet()
	for item := range s.Set {
		if !s2.Contains(item) {
			notIds.Add(item)
		}
	}
	return notIds
}

// MapInterfaceSet
type MapInterfaceSet struct {
	Sets map[string]*InterfaceSet
}

func NewMapInterfaceSet() *MapInterfaceSet {
	return &MapInterfaceSet{}
}

func (ms *MapInterfaceSet) check() {
	if ms.Sets == nil {
		ms.Sets = make(map[string]*InterfaceSet)
	}
}

func (ms *MapInterfaceSet) checkKey(key string) {
	ms.check()
	if ms.Sets[key] == nil {
		ms.Sets[key] = NewInterfaceSet()
	}
}

func (ms *MapInterfaceSet) Length() int {
	ms.check()
	return len(ms.Sets)
}

func (ms *MapInterfaceSet) Add(key string, inp interface{}) {
	ms.checkKey(key)
	ms.Sets[key].Add(inp)
}

func (ms *MapInterfaceSet) GetSlice(key string) []interface{} {
	ms.checkKey(key)
	if ms.Sets[key] == nil {
		return make([]interface{}, 0)
	}
	return ms.Sets[key].GetSlice()
}

func (ms *MapInterfaceSet) GetMap() map[string]*InterfaceSet {
	ms.check()
	return ms.Sets
}

func (ms *MapInterfaceSet) GetMapWithSlices() map[string][]interface{} {
	withSlices := make(map[string][]interface{})
	for key := range ms.Sets {
		withSlices[key] = ms.Sets[key].GetSlice()
	}
	return withSlices
}

func (ms *MapInterfaceSet) Contains(key, inp string) bool {
	ms.checkKey(key)
	return ms.GetMap()[key].Contains(inp)
}

func (ms *MapInterfaceSet) Unwind() *InterfaceSet {
	ms.check()
	items := NewInterfaceSet()
	for _, item := range ms.Sets {
		items.Update(item.GetSlice())
	}
	return items
}

func (ms *MapInterfaceSet) And(ms2 *MapInterfaceSet) *InterfaceSet {
	ms.check()
	intersection := NewInterfaceSet()
	var compareSet *InterfaceSet
	var iterateSet *InterfaceSet
	if ms2 == nil {
		return intersection
	}
	// Iterate the shorter set.
	if ms2.Length() > ms.Length() {
		compareSet = ms2.Unwind()
		iterateSet = ms.Unwind()
	} else {
		compareSet = ms.Unwind()
		iterateSet = ms2.Unwind()
	}
	for _, item := range iterateSet.GetSlice() {
		if compareSet.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s *MapInterfaceSet) Not(s2 *MapInterfaceSet) *InterfaceSet {
	s.check()
	compareA := s.Unwind()
	if s2 == nil {
		return compareA
	}
	notIds := NewInterfaceSet()
	compareB := s2.Unwind()
	for item := range compareA.Set {
		if !compareB.Contains(item) {
			notIds.Add(item)
		}
	}
	return notIds
}
