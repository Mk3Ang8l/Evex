export namespace models {
	
	export class Project {
	    id: string;
	    name: string;
	    description: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Section {
	    id: string;
	    project_id: string;
	    title: string;
	    content: string;
	    order: number;
	    tags: string[];
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Section(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.project_id = source["project_id"];
	        this.title = source["title"];
	        this.content = source["content"];
	        this.order = source["order"];
	        this.tags = source["tags"];
	        this.created_at = source["created_at"];
	    }
	}
	export class Source {
	    id: string;
	    section_id: string;
	    title: string;
	    url: string;
	    snippet: string;
	    content: string;
	    tags: string;
	    notes: string;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Source(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.section_id = source["section_id"];
	        this.title = source["title"];
	        this.url = source["url"];
	        this.snippet = source["snippet"];
	        this.content = source["content"];
	        this.tags = source["tags"];
	        this.notes = source["notes"];
	        this.created_at = source["created_at"];
	    }
	}

}

