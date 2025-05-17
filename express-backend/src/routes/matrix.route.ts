import { Router } from 'express';
import { analyzeMatrix } from '../controllers/matrix.controller';
import validateJWT from "../middlewares/validateJWT";

const router = Router();

router.post('/analyze', validateJWT, analyzeMatrix);

export default router;
