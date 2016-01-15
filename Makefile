export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

LIBDIR=${DESTDIR}/lib/systemd/system
BINDIR=${DESTDIR}/usr/lib/docker/

all:
	go build  -o registryacl .

install:
	install -d -m 0755 ${LIBDIR}
	install -m 644 systemd/registryacl-plugin.service ${LIBDIR}
	install -d -m 0755 ${BINDIR}
	install -m 755 registryacl ${BINDIR}

clean:
	rm -f registryacl
