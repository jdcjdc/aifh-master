/*
 * Artificial Intelligence for Humans
 * Volume 2: Nature Inspired Algorithms
 * Java Version
 * http://www.aifh.org
 * http://www.jeffheaton.com
 *
 * Code repository:
 * https://github.com/jeffheaton/aifh
 *
 * Copyright 2014 by Jeff Heaton
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
package com.heatonresearch.aifh.general.fns

/**
 * The Gaussian function is a Radial Basis Function that describes the typical "bell curve", or "normal distribution".
 * <p/>
 * The Gaussian function requires paramaters that specify the width (over all dimensions), as well as the
 * centers of each dimension.  So a 3d Gaussian would have the parameters lined up as follows:
 * <p/>
 * params[0] = width (of all dimensions),
 * <p/>
 * params[1] = center of dimension 0,
 * <p/>
 * params[2] = center of dimension 1,
 * <p/>
 * params[3] = center of dimension 3
 * <p/>
 * http://en.wikipedia.org/wiki/Gaussian_function
 *
 * @param theDimensions The number of dimensions.
 * @param theParams     A vector to hold the parameters.
 * @param theIndex      The index into the params vector.  You can store multiple RBF's in a vector.
 */
class GaussianFunction(theDimensions: Int, theParams: Array[Double], theIndex: Int)
  extends AbstractRBF(theDimensions, theParams, theIndex) {

  override def evaluate(x: Array[Double]): Double = {
    var value = 0.0
    val width = getWidth.abs
    for(i <- 0 until getDimensions) {
      val center = this.getCenter(i)
      value += Math.pow(x(i) - center, 2) / (2.0 * width * width)
    }
    Math.exp(-value)
  }
}