import { MongoClient } from "mongodb"
import { MONGO_URI } from '$env/static/private';

const client = new MongoClient(MONGO_URI)
await client.connect()
export default client.db('events') // select database