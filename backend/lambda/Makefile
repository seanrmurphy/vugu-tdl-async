SUBDIRS := onconnect ondisconnect default create-todo list-todos get-todo delete-todo update-todo

all: $(SUBDIRS)

$(SUBDIRS):
	make -C $@

.PHONY: all $(SUBDIRS)
