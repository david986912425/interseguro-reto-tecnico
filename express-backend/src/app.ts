import express from 'express';
import MatrixRoute from './routes/matrix.route';
import errorMiddleware from './middlewares/error.middleware';

const app = express();

app.use(express.json());
app.use('/api/matrix', MatrixRoute);
app.use(errorMiddleware);

export default app;
