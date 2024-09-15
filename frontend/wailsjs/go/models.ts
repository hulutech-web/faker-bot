export namespace amusement {
	
	export class Video {
	    item_cover: string;
	    share_url: string;
	    title: string;
	
	    static createFrom(source: any = {}) {
	        return new Video(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.item_cover = source["item_cover"];
	        this.share_url = source["share_url"];
	        this.title = source["title"];
	    }
	}
	export class ListItem {
	    rank_change: string;
	    video_list: Video[];
	    avatar: string;
	    effect_value: number;
	    follower_count: number;
	    nickname: string;
	    onbillbaord_times: number;
	    rank: number;
	
	    static createFrom(source: any = {}) {
	        return new ListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rank_change = source["rank_change"];
	        this.video_list = this.convertValues(source["video_list"], Video);
	        this.avatar = source["avatar"];
	        this.effect_value = source["effect_value"];
	        this.follower_count = source["follower_count"];
	        this.nickname = source["nickname"];
	        this.onbillbaord_times = source["onbillbaord_times"];
	        this.rank = source["rank"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Data {
	    list: ListItem[];
	    error_code: number;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.list = this.convertValues(source["list"], ListItem);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Extra {
	    error_code: number;
	    description: string;
	    sub_error_code: number;
	    sub_description: string;
	    logid: string;
	    now: number;
	
	    static createFrom(source: any = {}) {
	        return new Extra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	        this.sub_error_code = source["sub_error_code"];
	        this.sub_description = source["sub_description"];
	        this.logid = source["logid"];
	        this.now = source["now"];
	    }
	}
	
	export class Root {
	    data: Data;
	    extra: Extra;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], Data);
	        this.extra = this.convertValues(source["extra"], Extra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace drama {
	
	export class Video {
	    item_cover: string;
	    share_url: string;
	    title: string;
	
	    static createFrom(source: any = {}) {
	        return new Video(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.item_cover = source["item_cover"];
	        this.share_url = source["share_url"];
	        this.title = source["title"];
	    }
	}
	export class ListItem {
	    onbillbaord_times: number;
	    rank: number;
	    rank_change: string;
	    video_list: Video[];
	    avatar: string;
	    effect_value: number;
	    follower_count: number;
	    nickname: string;
	
	    static createFrom(source: any = {}) {
	        return new ListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.onbillbaord_times = source["onbillbaord_times"];
	        this.rank = source["rank"];
	        this.rank_change = source["rank_change"];
	        this.video_list = this.convertValues(source["video_list"], Video);
	        this.avatar = source["avatar"];
	        this.effect_value = source["effect_value"];
	        this.follower_count = source["follower_count"];
	        this.nickname = source["nickname"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Data {
	    list: ListItem[];
	    error_code: number;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.list = this.convertValues(source["list"], ListItem);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Extra {
	    error_code: number;
	    description: string;
	    sub_error_code: number;
	    sub_description: string;
	    logid: string;
	    now: number;
	
	    static createFrom(source: any = {}) {
	        return new Extra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	        this.sub_error_code = source["sub_error_code"];
	        this.sub_description = source["sub_description"];
	        this.logid = source["logid"];
	        this.now = source["now"];
	    }
	}
	
	export class Root {
	    data: Data;
	    extra: Extra;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], Data);
	        this.extra = this.convertValues(source["extra"], Extra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace entities {
	
	export class Param {
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    // Go type: sql
	    deleted_at: any;
	    url: string;
	    method: string;
	    body: string;
	    header: string;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Param(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.deleted_at = this.convertValues(source["deleted_at"], null);
	        this.url = source["url"];
	        this.method = source["method"];
	        this.body = source["body"];
	        this.header = source["header"];
	        this.label = source["label"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace hotsearch {
	
	export class Statistics {
	    play_count: number;
	    share_count: number;
	    comment_count: number;
	    digg_count: number;
	    download_count: number;
	    forward_count: number;
	
	    static createFrom(source: any = {}) {
	        return new Statistics(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.play_count = source["play_count"];
	        this.share_count = source["share_count"];
	        this.comment_count = source["comment_count"];
	        this.digg_count = source["digg_count"];
	        this.download_count = source["download_count"];
	        this.forward_count = source["forward_count"];
	    }
	}
	export class ListItem {
	    create_time: number;
	    is_reviewed: boolean;
	    share_url: string;
	    title: string;
	    video_status: number;
	    cover: string;
	    is_top: boolean;
	    item_id: string;
	    statistics: Statistics;
	
	    static createFrom(source: any = {}) {
	        return new ListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.create_time = source["create_time"];
	        this.is_reviewed = source["is_reviewed"];
	        this.share_url = source["share_url"];
	        this.title = source["title"];
	        this.video_status = source["video_status"];
	        this.cover = source["cover"];
	        this.is_top = source["is_top"];
	        this.item_id = source["item_id"];
	        this.statistics = this.convertValues(source["statistics"], Statistics);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Data {
	    description: string;
	    error_code: number;
	    list: ListItem[];
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.error_code = source["error_code"];
	        this.list = this.convertValues(source["list"], ListItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Extra {
	    error_code: number;
	    description: string;
	    sub_error_code: number;
	    sub_description: string;
	    logid: string;
	    now: number;
	
	    static createFrom(source: any = {}) {
	        return new Extra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	        this.sub_error_code = source["sub_error_code"];
	        this.sub_description = source["sub_description"];
	        this.logid = source["logid"];
	        this.now = source["now"];
	    }
	}
	
	export class Root {
	    data: Data;
	    extra: Extra;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], Data);
	        this.extra = this.convertValues(source["extra"], Extra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace hotsentence {
	
	export class Statistics {
	    comment_count: number;
	    digg_count: number;
	    download_count: number;
	    forward_count: number;
	    play_count: number;
	    share_count: number;
	
	    static createFrom(source: any = {}) {
	        return new Statistics(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.comment_count = source["comment_count"];
	        this.digg_count = source["digg_count"];
	        this.download_count = source["download_count"];
	        this.forward_count = source["forward_count"];
	        this.play_count = source["play_count"];
	        this.share_count = source["share_count"];
	    }
	}
	export class ListItem {
	    statistics: Statistics;
	    title: string;
	    cover: string;
	    create_time: number;
	    is_reviewed: boolean;
	    is_top: boolean;
	    item_id: string;
	    share_url: string;
	    video_status: number;
	
	    static createFrom(source: any = {}) {
	        return new ListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.statistics = this.convertValues(source["statistics"], Statistics);
	        this.title = source["title"];
	        this.cover = source["cover"];
	        this.create_time = source["create_time"];
	        this.is_reviewed = source["is_reviewed"];
	        this.is_top = source["is_top"];
	        this.item_id = source["item_id"];
	        this.share_url = source["share_url"];
	        this.video_status = source["video_status"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Data {
	    description: string;
	    error_code: number;
	    list: ListItem[];
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.error_code = source["error_code"];
	        this.list = this.convertValues(source["list"], ListItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Extra {
	    error_code: number;
	    description: string;
	    sub_error_code: number;
	    sub_description: string;
	    logid: string;
	    now: number;
	
	    static createFrom(source: any = {}) {
	        return new Extra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	        this.sub_error_code = source["sub_error_code"];
	        this.sub_description = source["sub_description"];
	        this.logid = source["logid"];
	        this.now = source["now"];
	    }
	}
	
	export class Root {
	    data: Data;
	    extra: Extra;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], Data);
	        this.extra = this.convertValues(source["extra"], Extra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace hotvideo {
	
	export class Extra {
	    error_code: number;
	    description: string;
	    sub_error_code: number;
	    sub_description: string;
	    logid: string;
	    now: number;
	
	    static createFrom(source: any = {}) {
	        return new Extra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	        this.sub_error_code = source["sub_error_code"];
	        this.sub_description = source["sub_description"];
	        this.logid = source["logid"];
	        this.now = source["now"];
	    }
	}
	export class  {
	    digg_count: number;
	    hot_value: number;
	    hot_words: string;
	    item_cover: string;
	    play_count: number;
	    rank: number;
	    share_url: string;
	    title: string;
	    author: string;
	    comment_count: number;
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.digg_count = source["digg_count"];
	        this.hot_value = source["hot_value"];
	        this.hot_words = source["hot_words"];
	        this.item_cover = source["item_cover"];
	        this.play_count = source["play_count"];
	        this.rank = source["rank"];
	        this.share_url = source["share_url"];
	        this.title = source["title"];
	        this.author = source["author"];
	        this.comment_count = source["comment_count"];
	    }
	}
	export class VideoData {
	    list: [];
	    error_code: number;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new VideoData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.list = this.convertValues(source["list"], );
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Root {
	    data: VideoData;
	    extra: Extra;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], VideoData);
	        this.extra = this.convertValues(source["extra"], Extra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace topic {
	
	export class Extra {
	    error_code: number;
	    description: string;
	    sub_error_code: number;
	    sub_description: string;
	    logid: string;
	    now: number;
	
	    static createFrom(source: any = {}) {
	        return new Extra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	        this.sub_error_code = source["sub_error_code"];
	        this.sub_description = source["sub_description"];
	        this.logid = source["logid"];
	        this.now = source["now"];
	    }
	}
	export class ListItem {
	    rank_change: string;
	    title: string;
	    effect_value: number;
	    rank: number;
	
	    static createFrom(source: any = {}) {
	        return new ListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rank_change = source["rank_change"];
	        this.title = source["title"];
	        this.effect_value = source["effect_value"];
	        this.rank = source["rank"];
	    }
	}
	export class VideoData {
	    list: ListItem[];
	    error_code: number;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new VideoData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.list = this.convertValues(source["list"], ListItem);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Root {
	    data: VideoData;
	    extra: Extra;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], VideoData);
	        this.extra = this.convertValues(source["extra"], Extra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace trending {
	
	export class ListItem {
	    hot_level: number;
	    label: number;
	    sentence: string;
	
	    static createFrom(source: any = {}) {
	        return new ListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hot_level = source["hot_level"];
	        this.label = source["label"];
	        this.sentence = source["sentence"];
	    }
	}
	export class Data {
	    has_more: boolean;
	    list: ListItem[];
	    total: number;
	    cursor: number;
	    description: string;
	    error_code: number;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.has_more = source["has_more"];
	        this.list = this.convertValues(source["list"], ListItem);
	        this.total = source["total"];
	        this.cursor = source["cursor"];
	        this.description = source["description"];
	        this.error_code = source["error_code"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Extra {
	    error_code: number;
	    description: string;
	    sub_error_code: number;
	    sub_description: string;
	    logid: string;
	    now: number;
	
	    static createFrom(source: any = {}) {
	        return new Extra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error_code = source["error_code"];
	        this.description = source["description"];
	        this.sub_error_code = source["sub_error_code"];
	        this.sub_description = source["sub_description"];
	        this.logid = source["logid"];
	        this.now = source["now"];
	    }
	}
	
	export class Root {
	    data: Data;
	    extra: Extra;
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], Data);
	        this.extra = this.convertValues(source["extra"], Extra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

