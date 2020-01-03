/*
 * Artificial Intelligence for Humans
 * Volume 1: Fundamental Algorithms
 * Scala Version
 * http://www.aifh.org
 * http://www.jeffheaton.com
 *
 * Code repository:
 * https://github.com/jeffheaton/aifh

 * Copyright 2013 by Jeff Heaton
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For more information on Heaton Research copyrights, licenses
 * and trademarks visit:
 * http://www.heatonresearch.com/copyright
 */
package com.heatonresearch.aifh.distance

/**
 * The Manhattan Distance (also known as Taxicab distance) is a Distance Metric used in machine learning.
 * This distance is used to compare how similar two vectors of uniform length are. A lower length indicates that
 * the two vectors are more similar than two vectors with a larger length.
 * <p/>
 * http://www.heatonresearch.com/wiki/Manhattan_Distance
 */
class ManhattanDistance extends CalculateDistance {
  override def calculate(position1: Vector[Double], position2: Vector[Double]): Double = {
    position1.zip(position2).map(v => Math.abs(v._1 - v._2)).sum
  }
}