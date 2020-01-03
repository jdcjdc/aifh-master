package tools;

/**
 * Created by jdeco on 19.04.17.
 */

//    public Map<String, Integer> encodeEquilateral(final int column, final double offValue, final double onValue) {
//        // remember the column name
//        final String name = this.headers[column];
//
//        // make space for it
//        final Map<String, Integer> classes = enumerateClasses(column);
//        final int classCount = classes.size();
//        // we'll use classCount - 1 columns, and we have the original column to reuse
//        insertColumns(column + 1, classCount - 2);
//
//        // perform the equilateral
//        final Equilateral eq = new Equilateral(classCount, offValue, onValue);
//
//        for (final Object[] obj : this.data) {
//            final int index = classes.get(obj[column].toString());
//
//            final double[] encoded = eq.encode(index);
//
//            for (int i = 0; i < classCount - 1; i++) {
//                obj[column + i] = encoded[i];
//            }
//        }

var headers []string

func SetHeaders(headers []string) {
	copy(headers, headers)
}



func EncodeEquilateral(column int, offValue float64, onValue float64) (myMap map[string]int)  {
	// remember the column name
	//final String name = this.headers[column];
	//
	//// make space for it
	//final Map<String, Integer> classes = enumerateClasses(column);
	//final int classCount = classes.size();
	//// we'll use classCount - 1 columns, and we have the original column to reuse
	//insertColumns(column + 1, classCount - 2);
	//
	//// perform the equilateral
	//final Equilateral eq = new Equilateral(classCount, offValue, onValue);
	//
	//for (final Object[] obj : this.data) {
	//    final int index = classes.get(obj[column].toString());
	//
	//    final double[] encoded = eq.encode(index);
	//
	//    for (int i = 0; i < classCount - 1; i++) {
	//        obj[column + i] = encoded[i];
	//    }
	return nil
}
