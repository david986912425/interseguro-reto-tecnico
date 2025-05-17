import { Request, Response, NextFunction } from 'express';
import { analyzeMatrices } from '../services/matrix.service';

export const analyzeMatrix = (req: Request, res: Response, next: NextFunction): void => {
    try {
        const { Q, R, rotated } = req.body;
        const result = analyzeMatrices(Q, R, rotated);
        res.json(result);
    } catch (error) {
        next(error);
    }
};
