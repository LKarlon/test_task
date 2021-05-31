package repository

const (

	qGetInfo = `
	select
    sv.name,
    bw.value,
    bw.interface_name,
	from servers sv
         join bandwidths bw on sv.id = bw.server_id
where sv.id=$1 ;`

	qDeleteInfo = `
	delete from bandwidths
	where server_id=$1 ;
	delete from servers
	where id=$1 ;`
)
