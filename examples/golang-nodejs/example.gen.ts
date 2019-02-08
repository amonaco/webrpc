
/* tslint:disable */

// This file has been generated by https://github.com/webrpc/webrpc
// Do not edit.


export enum Kind {
  Kind_USER = <number>1,
  Kind_ADMIN = <number>2
}
export const Kind_name = {
  '1': 'USER',
  '2': 'ADMIN'
}

export interface IEmpty {
  toJSON?(): object
}

export class Empty {
  private _data: object
  constructor(_data?: IEmpty) {
    this._data = {}
    if (_data) {
      
    }
  }
  
  public toJSON(): object {
    return this._data
  }
}

export interface IGetUserRequest {
  userID: number
  toJSON?(): object
}

export class GetUserRequest {
  private _data: object
  constructor(_data?: IGetUserRequest) {
    this._data = {}
    if (_data) {
      this._data['userID'] = _data['userID']!
      
    }
  }
  public get UserID(): number {
    return this._data['userID']!
  }
  public set UserID(value: number) {
    this._data['userID'] = value
  }
  
  public toJSON(): object {
    return this._data
  }
}

export interface IUser {
  id: number
  USERNAME: string
  created_at?: string
  toJSON?(): object
}

export class User {
  private _data: object
  constructor(_data?: IUser) {
    this._data = {}
    if (_data) {
      this._data['id'] = _data['id']!
      this._data['USERNAME'] = _data['USERNAME']!
      this._data['created_at'] = _data['created_at']!
      
    }
  }
  public get ID(): number {
    return this._data['id']!
  }
  public set ID(value: number) {
    this._data['id'] = value
  }
  public get Username(): string {
    return this._data['USERNAME']!
  }
  public set Username(value: string) {
    this._data['USERNAME'] = value
  }
  public get CreatedAt(): string {
    return this._data['created_at']!
  }
  public set CreatedAt(value: string) {
    this._data['created_at'] = value
  }
  
  public toJSON(): object {
    return this._data
  }
}

export interface IRandomStuff {
  meta: Map<string,any>
  metaNestedExample: Map<string,Map<string,number>>
  namesList: Array<string>
  numsList: Array<number>
  doubleArray: Array<Array<string>>
  listOfMaps: Array<Map<string,number>>
  listOfUsers: Array<IUser>
  mapOfUsers: Map<string,IUser>
  user: IUser
  toJSON?(): object
}

export class RandomStuff {
  private _data: object
  constructor(_data?: IRandomStuff) {
    this._data = {}
    if (_data) {
      this._data['meta'] = _data['meta']!
      this._data['metaNestedExample'] = _data['metaNestedExample']!
      this._data['namesList'] = _data['namesList']!
      this._data['numsList'] = _data['numsList']!
      this._data['doubleArray'] = _data['doubleArray']!
      this._data['listOfMaps'] = _data['listOfMaps']!
      this._data['listOfUsers'] = _data['listOfUsers']!
      this._data['mapOfUsers'] = _data['mapOfUsers']!
      this._data['user'] = _data['user']!
      
    }
  }
  public get Meta(): Map<string,any> {
    return this._data['meta']!
  }
  public set Meta(value: Map<string,any>) {
    this._data['meta'] = value
  }
  public get MetaNestedExample(): Map<string,Map<string,number>> {
    return this._data['metaNestedExample']!
  }
  public set MetaNestedExample(value: Map<string,Map<string,number>>) {
    this._data['metaNestedExample'] = value
  }
  public get NamesList(): Array<string> {
    return this._data['namesList']!
  }
  public set NamesList(value: Array<string>) {
    this._data['namesList'] = value
  }
  public get NumsList(): Array<number> {
    return this._data['numsList']!
  }
  public set NumsList(value: Array<number>) {
    this._data['numsList'] = value
  }
  public get DoubleArray(): Array<Array<string>> {
    return this._data['doubleArray']!
  }
  public set DoubleArray(value: Array<Array<string>>) {
    this._data['doubleArray'] = value
  }
  public get ListOfMaps(): Array<Map<string,number>> {
    return this._data['listOfMaps']!
  }
  public set ListOfMaps(value: Array<Map<string,number>>) {
    this._data['listOfMaps'] = value
  }
  public get ListOfUsers(): Array<IUser> {
    return this._data['listOfUsers']!
  }
  public set ListOfUsers(value: Array<IUser>) {
    this._data['listOfUsers'] = value
  }
  public get MapOfUsers(): Map<string,IUser> {
    return this._data['mapOfUsers']!
  }
  public set MapOfUsers(value: Map<string,IUser>) {
    this._data['mapOfUsers'] = value
  }
  public get User(): IUser {
    return this._data['user']!
  }
  public set User(value: IUser) {
    this._data['user'] = value
  }
  
  public toJSON(): object {
    return this._data
  }
}

export interface IExampleServiceService {
  Ping(headers: object): Promise<boolean>
  GetUser(params: IGetUserRequest, headers: object): Promise<User>
  
}


// Client

const ExampleServicePathPrefix = "/rpc/ExampleService/"

export class ExampleService implements IExampleServiceService {
  private hostname: string
  private fetch: Fetch
  private path = '/rpc/ExampleService/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = fetch
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }

  
  Ping(headers: object = {}): Promise<boolean> {
    return this.fetch(
      this.url('Ping'),
      
      createHTTPRequest({}, headers)
      
    ).then((res) => {
      if (!res.ok) {
        return throwHTTPError(res)
      }
      
      return res.json().then((_data) => {return <boolean>(_data)})
      
    })
  }
  
  GetUser(params: IGetUserRequest, headers: object = {}): Promise<User> {
    return this.fetch(
      this.url('GetUser'),
      
      createHTTPRequest(params, headers)
      
    ).then((res) => {
      if (!res.ok) {
        return throwHTTPError(res)
      }
      
      return res.json().then((_data) => {return new User(_data)})
      
    })
  }
  
}


// TODO: Server



export interface WebRPCError extends Error {
  code: string
  msg: string
	status: number
}

export const throwHTTPError = (resp: Response) => {
  return resp.json().then((err: WebRPCError) => { throw err })
}

export const createHTTPRequest = (body: object = {}, headers: object = {}): object => {
  return {
    method: 'POST',
    headers: { ...headers, 'Content-Type': 'application/json' },
    body: JSON.stringify(body || {})
  }
}

export type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>

