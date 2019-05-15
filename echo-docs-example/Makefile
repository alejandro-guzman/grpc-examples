.PHONY: proto-docs
proto-docs:
	docker run --rm \
		-v $(PWD)/docs:/out \
		-v $(PWD):/protos \
		pseudomuto/protoc-gen-doc --doc_opt=json,docs.json