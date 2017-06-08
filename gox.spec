# gox (loose) specification


## Mode Transitions
go -> gox-tag
go )-> pop (extra })

gox-tag -> go ({)
gox-tag )-> bare-words (>)
gox-tag )-> pop (/>)

bare-words -> go ({)
bare-words -> gox-tag (<(char))
bare-words )-> pop (</)