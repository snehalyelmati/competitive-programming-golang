Put:  1 1
insert node: &{<nil> <nil> 1 1}
this.capacity: 2
PUT.kv: map[1:0xc0000b8180]

Put:  2 2
insert node: &{<nil> <nil> 2 2}
this.capacity: 1
PUT.kv: map[1:0xc0000b8180 2:0xc0000b8200]

Get:  1
remove node: &{0xc0000b8200 0xc0000b8140 1 1}
this.capacity: 0
this.kv: map[2:0xc0000b8200]
insert node: &{0xc0000b8200 0xc0000b8140 1 1}
this.capacity: 1
GET.kv: map[1:0xc0000b8180 2:0xc0000b8200]

Put:  3 3
insert node: &{<nil> <nil> 3 3}
this.capacity: 0
remove node: &{0xc0000b8200 0xc0000b8160 0 0}
this.capacity: 0
this.kv: map[1:0xc0000b8180 2:0xc0000b8200]
PUT.kv: map[1:0xc0000b8180 2:0xc0000b8200 3:0xc0000b82e0]
Get:  2
remove node: &{0xc0000b8180 0xc0000b8160 2 2}
this.capacity: 0
this.kv: map[1:0xc0000b8180 3:0xc0000b82e0]
insert node: &{0xc0000b8180 0xc0000b8160 2 2}
this.capacity: 1
GET.kv: map[1:0xc0000b8180 2:0xc0000b8200 3:0xc0000b82e0]
Put:  4 4
insert node: &{<nil> <nil> 4 4}
this.capacity: 0
remove node: &{0xc0000b82e0 0xc0000b8160 1 1}
this.capacity: 0
this.kv: map[2:0xc0000b8200 3:0xc0000b82e0]
PUT.kv: map[2:0xc0000b8200 3:0xc0000b82e0 4:0xc0000b8400]
Get:  1
GET.kv: map[2:0xc0000b8200 3:0xc0000b82e0 4:0xc0000b8400]
Get:  3
remove node: &{0xc0000b8200 0xc0000b8160 3 3}
this.capacity: 0
this.kv: map[2:0xc0000b8200 4:0xc0000b8400]
insert node: &{0xc0000b8200 0xc0000b8160 3 3}
this.capacity: 1
GET.kv: map[2:0xc0000b8200 3:0xc0000b82e0 4:0xc0000b8400]
Get:  4
remove node: &{0xc0000b82e0 0xc0000b8200 4 4}
this.capacity: 0
this.kv: map[2:0xc0000b8200 3:0xc0000b82e0]
insert node: &{0xc0000b82e0 0xc0000b8200 4 4}
this.capacity: 1
GET.kv: map[2:0xc0000b8200 3:0xc0000b82e0 4:0xc0000b8400]
