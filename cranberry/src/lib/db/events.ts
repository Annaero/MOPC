import { MONGO_URI } from '$env/static/private';
import mongoose from 'mongoose';
mongoose.connect(MONGO_URI + "/events");
export default mongoose